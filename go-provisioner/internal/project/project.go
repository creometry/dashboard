package project

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

// Exportable functions

func ProvisionProjectNewUser(req ReqDataNewUser) (data RespDataProvisionProjectNewUser, err error) {
	// create gitRepo
	repoName, err := createGitRepo(req.GitRepoName, req.GitRepoUrl, req.GitRepoBranch)
	if err != nil {
		return RespDataProvisionProjectNewUser{}, err
	}
	fmt.Printf("Created repo : %s", repoName)

	// create rancher project
	projectId, err := createRancherProject(req.UsrProjectName, req.Plan)
	if err != nil {
		return RespDataProvisionProjectNewUser{}, err
	}
	// create random password
	password := generateRandomString(12)

	// create user
	userId, _, err := createUser(req.Username, password)
	if err != nil {
		return RespDataProvisionProjectNewUser{}, err
	}
	// add user to project
	_, err = AddUserToProject(userId, projectId)
	if err != nil {
		return RespDataProvisionProjectNewUser{}, err
	}

	// make post request to resources-service/namespace and pass the project name and id to create a namespace in the specific project
	nsName, err := createNamespace(req.UsrProjectName, projectId)

	fmt.Printf("Created namespace : %s", nsName)

	if err != nil {
		return RespDataProvisionProjectNewUser{}, err
	}
	//login as user to get token
	token, err := Login(userId, password)

	if err != nil {
		return RespDataProvisionProjectNewUser{}, err
	}

	resp := RespDataProvisionProjectNewUser{
		ProjectId: projectId,
		Token:     token,
		Password:  password,
	}
	return resp, nil

}

func ProvisionProject(req ReqData) (data RespDataProvisionProject, err error) {
	// create gitRepo
	repoName, err := createGitRepo(req.GitRepoName, req.GitRepoUrl, req.GitRepoBranch)
	if err != nil {
		return RespDataProvisionProject{}, err
	}
	fmt.Printf("Created repo : %s", repoName)

	// create rancher project
	projectId, err := createRancherProject(req.UsrProjectName, req.Plan)
	if err != nil {
		return RespDataProvisionProject{}, err
	}

	// add user to project
	_, err = AddUserToProject(req.UserId, projectId)
	if err != nil {
		return RespDataProvisionProject{}, err
	}

	// make post request to resources-service/namespace and pass the project name and id to create a namespace in the specific project
	nsName, err := createNamespace(req.UsrProjectName, projectId)

	fmt.Printf("Created namespace : %s", nsName)

	if err != nil {
		return RespDataProvisionProject{}, err
	}

	if err != nil {
		return RespDataProvisionProject{}, err
	}

	resp := RespDataProvisionProject{
		ProjectId: projectId,
	}
	return resp, nil

}

func GetNamespaceByAnnotation(annotations []string) (string, string, error) {

	// http get request to get the namespace list with http client
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s%s%s", os.Getenv("RANCHER_URL"), "/k8s/clusters/", os.Getenv("CLUSTER_ID"), "/v1/namespaces/"), nil)
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("RANCHER_TOKEN")))

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	// parse response body
	dt := RespDataNs{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}
	err = json.Unmarshal(body, &dt)
	if err != nil {
		return "", "", err
	}
	for _, annotation := range annotations {
		newAnnotation := fmt.Sprintf("%s:%s", os.Getenv("CLUSTER_ID"), strings.Split(annotation, ":")[0])
		for _, ns := range dt.Data {
			if ns.Metadata.Annotations["field.cattle.io/projectId"] == newAnnotation {
				return ns.Id, newAnnotation, nil
			}
		}
	}

	return "", "", nil

}

func GetKubeConfig(token string) (string, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v3/clusters/%s?action=generateKubeconfig", os.Getenv("RANCHER_URL"), os.Getenv("CLUSTER_ID")), nil)
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

func AddUserToProject(userId string, projectId string) (RespDataRoleBinding, error) {

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", os.Getenv("RANCHER_URL"), "/v3/projectroletemplatebindings"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"userId":"%s","projectId":"%s","roleTemplateId":"project-member"}`, userId, projectId))))
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
	dt := RespDataRoleBinding{}
	body, err := ioutil.ReadAll(resp.Body)
	log.Print(string(body))
	if err != nil {
		return RespDataRoleBinding{}, err
	}
	err = json.Unmarshal(body, &dt)
	if err != nil {
		return RespDataRoleBinding{}, err
	}
	return dt, nil

}

func GetUserByUsername(username string) (string, []string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s%s", os.Getenv("RANCHER_URL"), "/v3/users?username=", username), nil)
	if err != nil {
		return "", []string{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("RANCHER_TOKEN")))

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", []string{}, err
	}

	defer resp.Body.Close()

	dt := FindUserData{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", []string{}, err
	}
	err = json.Unmarshal(body, &dt)
	if err != nil {
		return "", []string{}, err
	}

	if len(dt.Data) == 0 {
		return "", []string{}, errors.New("user not found")
	}
	return dt.Data[0].Id, dt.Data[0].PrincipalIds, nil

}

// func FindUser(username string) (RespDataUser, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s%s", os.Getenv("RANCHER_URL"), "/v3/users?username=", username), nil)
// 	if err != nil {
// 		return RespDataUser{}, err
// 	}

// 	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("RANCHER_TOKEN")))

// 	client := &http.Client{}

// 	resp, err := client.Do(req)

// 	if err != nil {
// 		return RespDataUser{}, err
// 	}
// 	defer resp.Body.Close()

// 	// parse response body
// 	dt := UserData{}
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return RespDataUser{}, err
// 	}
// 	err = json.Unmarshal(body, &dt)
// 	if err != nil {
// 		return RespDataUser{}, err
// 	}

// 	// if user exists, login and return token
// 	if len(dt.Data) > 0 {
// 		token, err := loginAsUser(username, "testtesttest")
// 		if err != nil {
// 			return RespDataUser{}, err
// 		}

// 		// get his projectName
// 		pr, err := getProjectsOfUser(dt.Data[0].Id, dt.Data[0].PrincipalIds)
// 		if err != nil {
// 			return RespDataUser{}, err
// 		}

// 		// get namespace of project
// 		if len(pr) > 0 {
// 			rs, prId, err := GetNamespaceByAnnotation(pr)
// 			if err != nil {
// 				return RespDataUser{}, err
// 			}

// 			log.Printf("rs: %s", rs)
// 			log.Printf("prId: %s", prId)
// 			return RespDataUser{
// 				Id:        dt.Data[0].Id,
// 				Token:     token,
// 				Namespace: rs,
// 				ProjectId: strings.Split(prId, ":")[1],
// 			}, nil
// 		} else {
// 			return RespDataUser{
// 				Id:        dt.Data[0].Id,
// 				Token:     token,
// 				Namespace: "",
// 				ProjectId: "",
// 			}, nil
// 		}

// 	}

// 	// if user does not exist, create user and return token
// 	id, _, err := createUser(username)
// 	if err != nil {
// 		return RespDataUser{}, err
// 	}

// 	token, err := loginAsUser(username, "testtesttest")
// 	if err != nil {
// 		return RespDataUser{}, err
// 	}
// 	return RespDataUser{
// 		Id:        id,
// 		Token:     token,
// 		Namespace: "",
// 	}, nil
// }

func Login(username string, password string) (string, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", os.Getenv("RANCHER_URL"), "/v3-public/localProviders/local?action=login"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"username":"%s","password":"%s"}`, username, password))))
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

// Local functions

func createRancherProject(usrProjectName string, plan string) (string, error) {
	resourceQuota := genResourceQuotaFromPlan(plan)
	if resourceQuota == "nil" {
		return "", fmt.Errorf("invalid plan")
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", os.Getenv("RANCHER_URL"), "/v3/projects"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"name":"%s","clusterId":"%s",%s}`, usrProjectName, os.Getenv("CLUSTER_ID"), resourceQuota))))
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

func generateRandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func createUser(username string, password string) (string, []string, error) {

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", os.Getenv("RANCHER_URL"), "/v3/users"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"username":"%s","mustChangePassword": true,"password": "%s","principalIds": [ ]}`, username, password))))
	if err != nil {
		return "", []string{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("RANCHER_TOKEN")))

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", []string{}, err
	}

	defer resp.Body.Close()

	// parse response body
	dt := RespDataCreateUser{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", []string{}, err
	}
	err = json.Unmarshal(body, &dt)
	if err != nil {
		return "", []string{}, err
	}

	return dt.Id, dt.PrincipalIds, nil
}

func createGitRepo(name string, url string, branch string) (string, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/k8s/clusters/%s/v1/catalog.cattle.io.clusterrepos", os.Getenv("RANCHER_URL"), os.Getenv("CLUSTER_ID")), bytes.NewBuffer([]byte(fmt.Sprintf(`{
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
	  }`, name, url, branch))))

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
	dt := RespDataCreateGitRepo{}
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

func getProjectsOfUser(userId string, principalIds []string) ([]string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s%s", os.Getenv("RANCHER_URL"), "/v3/projectroletemplatebindings?userId=", userId), nil)
	if err != nil {
		return []string{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("RANCHER_TOKEN")))

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	// parse response body
	dt := RespDataProjectsByUser{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}
	err = json.Unmarshal(body, &dt)
	if err != nil {
		return []string{}, err
	}

	log.Println(dt)

	if len(dt.Data) > 0 {
		// return all the ids
		res := []string{}
		for _, v := range dt.Data {
			res = append(res, v.Id)
		}
		return res, nil

	}

	return []string{}, nil
}

func createNamespace(projectName string, projectId string) (string, error) {

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/k8s/clusters/%s/v1/namespaces/", os.Getenv("RANCHER_URL"), os.Getenv("CLUSTER_ID")), bytes.NewBuffer([]byte(fmt.Sprintf(`{"projectName":"%s","projectId":"%s"}`, projectName, projectId))))

	if err != nil {
		return "", err
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	// parse response body

	dt := CreateNsRespData{}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &dt)

	if err != nil {
		return "", err
	}

	if dt.Error != "" {
		return "", errors.New(dt.Error)
	}

	return dt.NsName, nil

}
