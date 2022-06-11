package project

import (
	"bytes"
	"context"
	_ "context"
	"crypto/sha1"
	_ "crypto/sha1"
	"encoding/base64"
	_ "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	_ "strings"
	"time"
	_ "time"

	"github.com/Creometry/dashboard/auth"
	_ "github.com/Creometry/dashboard/auth"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateProject(req ReqData) (kubeconfig string, err error) {
	// decode the id_token (JWT)
	

	// create rancher project
	projectId, err := createRancherProject(req.UsrProjectName)
	if err != nil {
		return "", err
	}

	// create user in rancher and get user id
	userId,principalIds, err := createUser(req.Username)

	if err != nil {
		return "", err
	}

	if len(principalIds) == 0 {
		return "", fmt.Errorf("User already exists")
	}

	// add user to project
	_, err = addUserToProject(userId, principalIds,projectId)
	if err != nil {
		return "", err
	}

	// create a new namespace with annotation "projectId"
	nsClient := auth.MyClientSet.CoreV1().Namespaces()

	// create a random hash and append it to the namespace name
	h := sha1.New()
	h.Write([]byte(time.Now().String()))
	b := h.Sum(nil)
	rand := base64.URLEncoding.EncodeToString(b)
	// delete every special character in the random hash
	rand = strings.Replace(rand, "=", "", -1)
	nsName := req.Namespace + "-" + strings.ToLower(rand+"x")


	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: 	nsName,
			Annotations: map[string]string{
				"field.cattle.io/projectId": fmt.Sprintf(projectId),
			},
			Labels: map[string]string{
				"field.cattle.io/projectId": strings.Split(projectId, ":")[1],
			},
		},
	}

	newNs, err := nsClient.Create(context.TODO(), ns, metav1.CreateOptions{})
	if err != nil {
		return "", err
	}
	log.Println("Created Namespace:", newNs.Name)

	// create a new resource quota based on the plan
	quotaClient := auth.MyClientSet.CoreV1().ResourceQuotas(newNs.Name)
	var quota *v1.ResourceQuota
	switch req.Plan {
	case "Starter":
		quota = createResourseQuota("1", "1Gi", newNs.Name)
	case "Dev":
		quota = createResourseQuota("2", "2Gi", newNs.Name)
	case "Pro":
		quota = createResourseQuota("4", "4Gi", newNs.Name)
	}

	newQuota, err := quotaClient.Create(context.TODO(), quota, metav1.CreateOptions{})

	if err != nil {
		return "", err
	}
	log.Println("Created Quota:", newQuota.Name)

	// login as user to get token
	/* token, err := loginAsUser(req.Username, "testtesttest")

	if err != nil {
		return "", err
	} */


	// get kubeconfig (still not working)
	/*kubeconfig, err = getKubeConfig(req.Id_token,projectId)
	if err != nil {
		return "", err
	}*/
	return "kubeconfig", nil

}

func createRancherProject(usrProjectName string) (string, error) {
	req, err := http.NewRequest("POST", os.Getenv("CREATE_PROJECT_URL"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"name":"%s","clusterId":"%s"}`, usrProjectName, os.Getenv("CLUSTER_ID")))))
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

func createResourseQuota(cpu string, memory string, namespace string) *v1.ResourceQuota {
	return &v1.ResourceQuota{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-quota", namespace),
			Namespace: namespace,
		},
		Spec: v1.ResourceQuotaSpec{
			Hard: v1.ResourceList{
				v1.ResourceLimitsCPU:    resource.MustParse(cpu),
				v1.ResourceLimitsMemory: resource.MustParse(memory),
			},
		},
	}
}

func getKubeConfig(token string,projectId string)(string, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("https://tn.cloud.creometry.com/v3/projects/%s?action=generateKubeconfig",projectId), nil)
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

	req, err := http.NewRequest("POST", os.Getenv("ADD_USER_TO_PROJECT_URL"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"userId":"%s/%s","projectId":"%s","roleTemplateId":"project-member"}`, userId, principalIds[0],projectId))))
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
	req, err := http.NewRequest("POST", os.Getenv("LOGIN_USER_URL"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"username":%s,"password":%s,"description":null,"ttl":0,"responseType":null}`, username,password))))
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