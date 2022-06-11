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

	"github.com/Creometry/dashboard/auth"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateProject(req ReqData) (kubeconfig string, err error) {
	// decode the id token base64 encoded
	decoded, err := base64.RawURLEncoding.DecodeString(req.Id_token)
	if err != nil {
		return "", err
	}

	// get decoded data as string
	decodedString := string(decoded)
	log.Println(decodedString)

	// create rancher project
	projectId, err := createRancherProject(req.UsrProjectName)
	if err != nil {
		return "", err
	}

	// create user in rancher and get user id
	userId, err := createUser(req.Username)

	if err != nil {
		return "", err
	}

	// add user to project
	_, err = addUserToProject(userId, projectId)
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

	nsName := req.Namespace + "-" + rand



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
	kubeconfig, err = getKubeConfig(req.Id_token,projectId)
	if err != nil {
		return "", err
	}
	return kubeconfig, nil

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


func addUserToProject(userId string,projectId string) (RespDataRoleBinding, error) {

	req, err := http.NewRequest("POST", os.Getenv("ADD_USER_TO_PROJECT_URL"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"userId":"%s","projectId":"%s","roleTemplateId":"project-member"}`, userId, projectId))))
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
	if err != nil {
		return RespDataRoleBinding{}, err
	}
	err = json.Unmarshal(body, &dt)
	if err != nil {
		return RespDataRoleBinding{}, err
	}

	return dt, nil

}

func createUser(username string)(string,error){
	req, err := http.NewRequest("POST", os.Getenv("CREATE_USER_URL"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"username":%s,"description": "",
	"mustChangePassword": false,
	"name": "%s",
	"password": "testtesttest",
	"principalIds": [ ],}`, username,username))))
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
	dt := RespDataCreateUser{}
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