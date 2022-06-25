package project

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Creometry/dashboard/auth"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateProject(req ReqData) (data RespDataCreateProjectAndRepo, err error) {
	// decode the id_token (JWT)
	
	// create gitRepo
	repoName,err:= createGitRepo(req.GitRepoName,req.GitRepoUrl,req.GitRepoBranch)
	if err != nil {
		return RespDataCreateProjectAndRepo{}, err
	}
	fmt.Printf("Created repo : %s",repoName)

	// create rancher project
	projectId, err := createRancherProject(req.UsrProjectName,req.Plan)
	if err != nil {
		return RespDataCreateProjectAndRepo{}, err
	}

	// create user in rancher and get user id
	userId,principalIds, err := createUser(req.Username)

	if err != nil {
		return RespDataCreateProjectAndRepo{}, err
	}

	if len(principalIds) == 0 {
		return RespDataCreateProjectAndRepo{}, fmt.Errorf("user already exists")
	}

	// add user to project
	_, err = addUserToProject(userId, principalIds,projectId)
	if err != nil {
		return RespDataCreateProjectAndRepo{}, err
	}

	// create a new namespace with annotation "projectId"
	nsClient := auth.MyClientSet.CoreV1().Namespaces()

	// create a random hash and append it to the namespace name
	h := sha1.New()
	h.Write([]byte(time.Now().String()))
	b := h.Sum(nil)
	rand := base64.URLEncoding.EncodeToString(b)
	// replace every special character in the random hash with a random letter
	rand = strings.Replace(rand, "+", "x", -1)
	rand = strings.Replace(rand, "/", "x", -1)
	rand = strings.Replace(rand, "=", "x", -1)
	rand = strings.Replace(rand, ".", "x", -1)
	rand = strings.Replace(rand, "-", "x", -1)
	rand = strings.Replace(rand, "_", "x", -1)
	rand = strings.Replace(rand, "*", "x", -1)
	rand = strings.Replace(rand, " ", "x", -1)
	rand = strings.Replace(rand, ",", "x", -1)

	nsName := strings.ToLower(req.UsrProjectName) + "-" + strings.ToLower(rand)


	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: 	nsName,
			Annotations: map[string]string{
				"field.cattle.io/projectId": projectId,
			},
			Labels: map[string]string{
				"field.cattle.io/projectId": strings.Split(projectId, ":")[1],
			},
		},
	}

	newNs, err := nsClient.Create(context.TODO(), ns, metav1.CreateOptions{})
	if err != nil {
		return RespDataCreateProjectAndRepo{}, err
	}
	log.Println("Created Namespace:", newNs.Name)


	// login as user to get token
	token, err := loginAsUser(req.Username, "testtesttest")

	if err != nil {
		return RespDataCreateProjectAndRepo{}, err
	} 

	resp :=RespDataCreateProjectAndRepo{
		User_token: token,
		Namespace: "newNs.Name",
	}
	return resp, nil

}

func createRancherProject(usrProjectName string,plan string) (string, error) {
	resourceQuota:= genResourceQuotaFromPlan(plan)
	if resourceQuota == "nil" {
		return "", fmt.Errorf("invalid plan")
	}
	req, err := http.NewRequest("POST", os.Getenv("CREATE_PROJECT_URL"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"name":"%s","clusterId":"%s",%s}`, usrProjectName, os.Getenv("CLUSTER_ID"),resourceQuota))))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("RANCHER_TOKEN")))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	// parse response body
	dt := RespData{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &dt)
	if err != nil {
		log.Fatal(err)
	}

	return dt.ProjectId, nil
}


func genResourceQuotaFromPlan(plan string) string {
	switch plan {
	case "Starter":
		return `"namespaceDefaultResourceQuota": {
			"limit": {
			"configMaps": "10",
			"limitsCpu": "1000m",
			"limitsMemory": "2000Mi",
			"persistentVolumeClaims": "10",
			"pods": "50",
			"replicationControllers": "15",
			"requestsStorage": "50000Mi",
			"secrets": "20",
			"services": "50",
			"servicesLoadBalancers": "0",
			"servicesNodePorts": "0"
			}
			},
			"resourceQuota": {
			"limit": {
			"configMaps": "10",
			"limitsCpu": "1000m",
			"limitsMemory": "2000Mi",
			"persistentVolumeClaims": "10",
			"pods": "100",
			"replicationControllers": "30",
			"requestsStorage": "50000Mi",
			"secrets": "20",
			"services": "50",
			"servicesLoadBalancers": "0",
			"servicesNodePorts": "0"
			},
			"usedLimit": { }
			}
		`
	case "Pro":
		return `"namespaceDefaultResourceQuota": {
			"limit": {
			"configMaps": "20",
			"limitsCpu": "2000m",
			"limitsMemory": "4000Mi",
			"persistentVolumeClaims": "20",
			"pods": "100",
			"replicationControllers": "25",
			"requestsStorage": "50000Mi",
			"secrets": "20",
			"services": "50",
			"servicesLoadBalancers": "0",
			"servicesNodePorts": "0"
			}
			},
			"resourceQuota": {
			"limit": {
			"configMaps": "20",
			"limitsCpu": "2000m",
			"limitsMemory": "4000Mi",
			"persistentVolumeClaims": "20",
			"pods": "100",
			"replicationControllers": "25",
			"requestsStorage": "50000Mi",
			"secrets": "20",
			"services": "50",
			"servicesLoadBalancers": "0",
			"servicesNodePorts": "0"
			},
			"usedLimit": { }
			}
		`
	case "Elite":
		return `"namespaceDefaultResourceQuota": {
			"limit": {
			"configMaps": "20",
			"limitsCpu": "4000m",
			"limitsMemory": "8000Mi",
			"persistentVolumeClaims": "30",
			"pods": "200",
			"replicationControllers": "50",
			"requestsStorage": "200000Mi",
			"secrets": "20",
			"services": "100",
			"servicesLoadBalancers": "0",
			"servicesNodePorts": "0"
			}
			},
			"resourceQuota": {
			"limit": {
			"configMaps": "20",
			"limitsCpu": "4000m",
			"limitsMemory": "8000Mi",
			"persistentVolumeClaims": "30",
			"pods": "200",
			"replicationControllers": "50",
			"requestsStorage": "200000Mi",
			"secrets": "20",
			"services": "100",
			"servicesLoadBalancers": "0",
			"servicesNodePorts": "0"
			},
			"usedLimit": { }
			}
		`
	}
	return "nil"
}

func GetKubeConfig(token string)(string, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("https://tn.cloud.creometry.com/v3/clusters/%s?action=generateKubeconfig",os.Getenv("CLUSTER_ID")), nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	// parse response body
	dt := Kubeconfig{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &dt)
	if err != nil {
		log.Fatal(err)
	}

	return dt.Config, nil
}


func addUserToProject(userId string,principalIds []string,projectId string) (RespDataRoleBinding, error) {

	req, err := http.NewRequest("POST", os.Getenv("ADD_USER_TO_PROJECT_URL"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"userId":"%s","projectId":"%s","roleTemplateId":"project-member"}`, userId,projectId))))
	if err != nil {
		return RespDataRoleBinding{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("RANCHER_TOKEN")))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return RespDataRoleBinding{}, err
	}

	defer resp.Body.Close()
	// parse response body
	dt:= RespDataRoleBinding{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return RespDataRoleBinding{}, err
	}
	err = json.Unmarshal(body, &dt)
	if err != nil {
		return RespDataRoleBinding{}, err
	}
	return dt, nil

}

//func addClusterRoleBinding

func createUser(username string)(string,[]string,error){
	req, err := http.NewRequest("POST", os.Getenv("CREATE_USER_URL"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"username":"%s","mustChangePassword": false,"password": "testtesttest","principalIds": [ ]}`, username))))
	if err != nil {
		return "",[]string{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("RANCHER_TOKEN")))

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "",[]string{}, err
	}
	
	defer resp.Body.Close()
	
	// parse response body
	dt:=RespDataCreateUser{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", []string{},err
	}
	err = json.Unmarshal(body, &dt)
	if err != nil {
		return "", []string{},err
	}

	return dt.Id,dt.PrincipalIds, nil

}

func loginAsUser(username string,password string )(string,error){
	req, err := http.NewRequest("POST", os.Getenv("LOGIN_USER_URL"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"username":"%s","password":"%s"}`, username,password))))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// parse response body
	dt := RespDataLogin{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &dt)
	if err != nil {
		return "", err
	}

	return dt.Token, nil

}

func createGitRepo(name string,url string,branch string)(string, error){
	req, err := http.NewRequest("POST", fmt.Sprintf("https://tn.cloud.creometry.com/k8s/clusters/%s/v1/catalog.cattle.io.clusterrepos",os.Getenv("CLUSTER_ID")), bytes.NewBuffer([]byte(fmt.Sprintf(`{
		"type": "catalog.cattle.io.clusterrepo",
		"metadata": {
		  "name": "%s"
		},
		"spec": {
		  "url": "",
		  "clientSecret": null,
		  "gitRepo": "%s",
		  "gitBranch": "%s"
		}
	  }`,name,url,branch))))

	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("RANCHER_TOKEN")))

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}
	
	defer resp.Body.Close()
	
	// parse response body
	dt:=RespDataCreateGitRepo{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &dt)
	if err != nil {
		return "", err
	}

	return dt.Id, nil
}

