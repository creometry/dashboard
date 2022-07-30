package project

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/Creometry/dashboard/go-provisioner/auth"
	"github.com/Seifbarouni/fast-utils/utils"
	"github.com/google/uuid"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Exportable functions

func ProvisionProject(req ReqData) (data RespDataProvisionProject, err error) {
	// check paymee payment
	_, err = checkPayment(req.PaymentToken)
	if err != nil {
		return RespDataProvisionProject{}, err
	}

	// create rancher project
	projectId, createdTS, p_uuid, err := createRancherProject(req.UsrProjectName, req.Plan)
	if err != nil {
		return RespDataProvisionProject{}, err
	}
	log.Println(p_uuid)

	// convert createdTS to time.Time
	t := time.Unix(0, createdTS)

	//create billing account
	if req.BillingAccountId == "1" {
		// create billing account
		accountId, err := createBillingAccount(req, projectId, t)
		if err != nil {
			return RespDataProvisionProject{}, err
		}
		log.Println(accountId)
	} else {
		// convert string to uuid
		uid := uuid.MustParse(req.BillingAccountId)
		// add project to billing account
		prId, err := addProjectToBillingAccount(uid, projectId, t, req.Plan)
		if err != nil {
			return RespDataProvisionProject{}, err
		}
		log.Println(prId)
	}

	// add user to project
	_, err = AddUserToProject(req.UserId, projectId)
	if err != nil {
		return RespDataProvisionProject{}, err
	}

	// create k8s namespace
	nsName, err := createNamespace(req.UsrProjectName, projectId)

	fmt.Printf("Created namespace : %s", nsName)

	if err != nil {
		return RespDataProvisionProject{}, err
	}

	// create gitRepo
	if req.GitRepoUrl != "" && req.GitRepoBranch != "" && req.GitRepoName != "" {
		repoName, err := createGitRepo(req.GitRepoName, req.GitRepoUrl, req.GitRepoBranch)
		if err != nil {
			return RespDataProvisionProject{}, err
		}
		fmt.Printf("Created repo : %s", repoName)
	}

	resp := RespDataProvisionProject{
		ProjectId: projectId,
	}
	return resp, nil

}

func GetNamespaceByAnnotation(annotations []string) (string, string, error) {

	clusterId, err := utils.GetVariable("config", "CLUSTER_ID")
	if err != nil {
		return "", "", err
	}

	rancherURL, err := utils.GetVariable("config", "RANCHER_URL")
	if err != nil {
		return "", "", err
	}

	rancherToken, err := utils.GetVariable("secrets", "RANCHER_TOKEN")
	if err != nil {
		return "", "", err
	}

	// http get request to get the namespace list with http client
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s%s%s", rancherURL, "/k8s/clusters/", clusterId, "/v1/namespaces/"), nil)
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", rancherToken))

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
		newAnnotation := fmt.Sprintf("%s:%s", clusterId, strings.Split(annotation, ":")[0])
		for _, ns := range dt.Data {
			if ns.Metadata.Annotations["field.cattle.io/projectId"] == newAnnotation {
				return ns.Id, newAnnotation, nil
			}
		}
	}

	return "", "", nil

}

func GetKubeConfig(token string) (string, error) {

	clusterId, err := utils.GetVariable("config", "CLUSTER_ID")
	if err != nil {
		return "", err
	}

	rancherURL, err := utils.GetVariable("config", "RANCHER_URL")
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v3/clusters/%s?action=generateKubeconfig", rancherURL, clusterId), nil)
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

	rancherURL, err := utils.GetVariable("config", "RANCHER_URL")
	if err != nil {
		return RespDataRoleBinding{}, err
	}

	rancherToken, err := utils.GetVariable("secrets", "RANCHER_TOKEN")
	if err != nil {
		return RespDataRoleBinding{}, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", rancherURL, "/v3/projectroletemplatebindings"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"userId":"%s","projectId":"%s","roleTemplateId":"project-member"}`, userId, projectId))))
	if err != nil {
		return RespDataRoleBinding{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", rancherToken))

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
	rancherURL, err := utils.GetVariable("config", "RANCHER_URL")
	if err != nil {
		return "", []string{}, err
	}

	rancherToken, err := utils.GetVariable("secrets", "RANCHER_TOKEN")
	if err != nil {
		return "", []string{}, err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s%s", rancherURL, "/v3/users?username=", username), nil)
	if err != nil {
		return "", []string{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", rancherToken))

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

func Login(username string, password string) (string, string, string, error) {

	rancherURL, err := utils.GetVariable("config", "RANCHER_URL")
	if err != nil {
		return "", "", "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", rancherURL, "/v3-public/localProviders/local?action=login"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"username":"%s","password":"%s"}`, username, password))))
	if err != nil {
		return "", "", "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", "", "", err
	}
	defer resp.Body.Close()

	// parse response body
	dt := RespDataLogin{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", err
	}
	err = json.Unmarshal(body, &dt)
	if err != nil {
		return "", "", "", err
	}

	return dt.Id, dt.Token, dt.UUID, nil

}

func Register(username string) (string, string, string, string, error) {

	rancherURL, err := utils.GetVariable("config", "RANCHER_URL")
	if err != nil {
		return "", "", "", "", err
	}

	rancherToken, err := utils.GetVariable("secrets", "RANCHER_TOKEN")
	if err != nil {
		return "", "", "", "", err
	}

	password := generateRandomString(16)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", rancherURL, "/v3/users"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"username":"%s","mustChangePassword": true,"password": "%s","enabled": true,"type":"user"}`, username, password))))
	if err != nil {
		return "", "", "", "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", rancherToken))

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", "", "", "", err
	}

	defer resp.Body.Close()

	// parse response body
	dt := RespDataCreateUser{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", "", err
	}
	err = json.Unmarshal(body, &dt)
	if err != nil {
		return "", "", "", "", err
	}

	err = createGlobalRoleBinding(dt.Id)

	if err != nil {
		return "", "", "", "", err
	}
	// login user
	id, token, uuid, err := Login(username, password)
	if err != nil {
		return "", "", "", "", err
	}

	return id, token, password, uuid, nil
}

// Local functions

func createRancherProject(usrProjectName string, plan string) (string, int64, string, error) {
	resourceQuota := genResourceQuotaFromPlan(plan)
	if resourceQuota == "nil" {
		return "", 0, "", fmt.Errorf("invalid plan")
	}

	clusterId, err := utils.GetVariable("config", "CLUSTER_ID")
	if err != nil {
		return "", 0, "", err
	}

	rancherURL, err := utils.GetVariable("config", "RANCHER_URL")
	if err != nil {
		return "", 0, "", err
	}

	rancherToken, err := utils.GetVariable("secrets", "RANCHER_TOKEN")
	if err != nil {
		return "", 0, "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", rancherURL, "/v3/projects"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"name":"%s","clusterId":"%s",%s}`, usrProjectName, clusterId, resourceQuota))))
	if err != nil {
		return "", 0, "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", rancherToken))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", 0, "", err
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

	return dt.ProjectId, dt.CreatedTS, dt.UUID, nil
}

func createGlobalRoleBinding(id string) error {

	rancherURL, err := utils.GetVariable("config", "RANCHER_URL")
	if err != nil {
		return err
	}

	rancherToken, err := utils.GetVariable("secrets", "RANCHER_TOKEN")
	if err != nil {
		return err
	}

	req2, err := http.NewRequest("POST", fmt.Sprintf("%s%s", rancherURL, "/v3/globalrolebindings"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"type":"globalRoleBinding","globalRoleId":"user","userId":"%s"}`, id))))

	if err != nil {
		return err
	}

	req2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", rancherToken))

	client2 := &http.Client{}

	resp2, err := client2.Do(req2)

	if err != nil {
		return err
	}

	defer resp2.Body.Close()

	// parse response body
	dt2 := RespDataCreateUser{}
	body2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body2, &dt2)
	if err != nil {
		return err
	}

	return nil
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

func createGitRepo(name string, url string, branch string) (string, error) {
	clusterId, err := utils.GetVariable("config", "CLUSTER_ID")
	if err != nil {
		return "", err
	}

	rancherURL, err := utils.GetVariable("config", "RANCHER_URL")
	if err != nil {
		return "", err
	}

	rancherToken, err := utils.GetVariable("secrets", "RANCHER_TOKEN")
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/k8s/clusters/%s/v1/catalog.cattle.io.clusterrepos", rancherURL, clusterId), bytes.NewBuffer([]byte(fmt.Sprintf(`{
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

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", rancherToken))

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
	rancherURL, err := utils.GetVariable("config", "RANCHER_URL")
	if err != nil {
		return []string{}, err
	}

	rancherToken, err := utils.GetVariable("secrets", "RANCHER_TOKEN")
	if err != nil {
		return []string{}, err
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s%s", rancherURL, "/v3/projectroletemplatebindings?userId=", userId), nil)
	if err != nil {
		return []string{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", rancherToken))

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

	nsClient := auth.MyClientSet.CoreV1().Namespaces()

	nsName := strings.ToLower(projectName) + "-" + generateRandomString(20)

	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: nsName,
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
		return "", err
	}
	return newNs.Name, nil
}

func checkPayment(token string) (CheckPaymeePaymentResponse, error) {
	paymeeURL, err := utils.GetVariable("config", "PAYMEE_URL")
	if err != nil {
		return CheckPaymeePaymentResponse{}, err
	}
	paymeeToken, err := utils.GetVariable("secrets", "PAYMEE_TOKEN")
	if err != nil {
		return CheckPaymeePaymentResponse{}, err
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/payments/%s/check", paymeeURL, token), nil)
	if err != nil {
		return CheckPaymeePaymentResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Token %s", paymeeToken))

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return CheckPaymeePaymentResponse{}, err
	}

	defer resp.Body.Close()

	// parse response body
	dt := CheckPaymeePaymentResponse{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return CheckPaymeePaymentResponse{}, err
	}
	err = json.Unmarshal(body, &dt)
	if err != nil {
		return CheckPaymeePaymentResponse{}, err
	}

	if dt.Message != "Success" || dt.Data.BuyerId == 0 {
		return CheckPaymeePaymentResponse{}, errors.New("payment failed")
	}

	return dt, nil
}

func createBillingAccount(req ReqData, projectId string, t time.Time) (string, error) {
	billingURL, err := utils.GetVariable("config", "BILLING_URL")
	if err != nil {
		return "", err
	}

	clId := strings.Split(projectId, ":")[0]
	prId := strings.Split(projectId, ":")[1]

	companyName := ""
	taxId := ""

	if req.IsCompany {
		companyName = req.CompanyName
		taxId = req.TaxId
	}
	reqBody := ReqDataCreateBillingAccount{
		Company: Company{
			IsCompany: req.IsCompany,
			TaxId:     taxId,
			Name:      companyName,
		},
		BillingAdmins: []Admin{
			{
				UUID:         req.UUID,
				Email:        req.Email,
				Phone_number: req.Phone,
			},
		},
		Projects: []Project{
			{
				ProjectId:         prId,
				ClusterId:         clId,
				CreationTimeStamp: t,
				State:             "active",
				Plan:              req.Plan,
			},
		},
	}


	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	r, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/CreateBillingAccount", billingURL), bytes.NewBuffer(reqBodyJson))

	if err != nil {
		return "", err
	}

	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(r)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != 201 {
		return "", errors.New("billing account creation failed")
	}

	defer resp.Body.Close()

	// parse response body
	dt := RespDataCreateBillingAccount{}

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

func addProjectToBillingAccount(billingAccountId uuid.UUID, projectId string, t time.Time, plan string) (string, error) {
	billingURL, err := utils.GetVariable("config", "BILLING_URL")
	if err != nil {
		return "", err
	}

	clId := strings.Split(projectId, ":")[0]
	prId := strings.Split(projectId, ":")[1]

	reqBody := &ReqDataAddProjectToBillingAccount{
		BillingAccountUUID: billingAccountId,
		ProjectId:          prId,
		ClusterId:          clId,
		CreationTimeStamp:  t,
		Plan:               plan,
		State:              "active",
	}

	b, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", billingURL, "/v1/addproject"), bytes.NewBuffer(b))

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
	dt := RespDataProvisionProject{}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &dt)

	if err != nil {
		return "", err
	}

	return dt.ProjectId, nil
}
