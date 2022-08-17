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
	"github.com/zemirco/keycloak"
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

func GetUserByUsername(username, region string) (string, []string, error) {
	// rancherURL, err := utils.GetVariable("config", "RANCHER_URL")
	// if err != nil {
	// 	return "", []string{}, err
	// }

	rancherToken, err := utils.GetVariable("secrets", "RANCHER_TOKEN")
	if err != nil {
		return "", []string{}, err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s%s", region, "/v3/users?username=", username), nil)
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

func Register(username, email string) error {

	// rancherURL, err := utils.GetVariable("config", "RANCHER_URL")
	// if err != nil {
	// 	return err
	// }

	// rancherToken, err := utils.GetVariable("secrets", "RANCHER_TOKEN")
	// if err != nil {
	// 	return err
	// }

	creometryGmail, err := utils.GetVariable("config", "CREOMETRY_GMAIL")
	if err != nil {
		return err
	}

	gmailPassword, err := utils.GetVariable("secrets", "GMAIL_PASSWORD")
	if err != nil {
		return err
	}

	password := generateRandomString(16)

	// req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", rancherURL, "/v3/users"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"username":"%s","mustChangePassword": true,"password": "%s","enabled": true,"type":"user"}`, username, password))))
	// if err != nil {
	// 	return err
	// }

	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", rancherToken))

	// client := &http.Client{}

	// resp, err := client.Do(req)

	// if err != nil {
	// 	return err
	// }

	// defer resp.Body.Close()

	// // parse response body
	// dt := RespDataCreateUser{}
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }
	// err = json.Unmarshal(body, &dt)
	// if err != nil {
	// 	return err
	// }

	// err = createGlobalRoleBinding(dt.Id)

	// if err != nil {
	// 	return err
	// }

	// create user in keycloak
	err = createKeyCloakUser(username, email, "", "", "creometry", password)

	if err != nil {
		return err
	}

	// send email
	err = utils.SendEmail(
		creometryGmail,
		email,
		// needs to be changed to the actual creometry gmail password
		gmailPassword,
		"Creometry Registration",
		fmt.Sprintf("Password: %s\nYou can use this password to log in to Creometry and Rancher dashboards.", password),
	)

	if err != nil {
		return err
	}

	return nil
}

func ResetPassword(userId, email, newPassword string) error {

	rancherURL, err := utils.GetVariable("config", "RANCHER_URL")
	if err != nil {
		return err
	}

	rancherToken, err := utils.GetVariable("secrets", "RANCHER_TOKEN")
	if err != nil {
		return err
	}

	creometryGmail, err := utils.GetVariable("config", "CREOMETRY_GMAIL")
	if err != nil {
		return err
	}

	gmailPassword, err := utils.GetVariable("secrets", "GMAIL_PASSWORD")
	if err != nil {
		return err
	}

	var password string
	if newPassword == "" {
		password = generateRandomString(16)
	} else {
		password = newPassword
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s?action=setpassword", rancherURL, "/v3/users/", userId), bytes.NewBuffer([]byte(fmt.Sprintf(`{"newPassword":"%s"}`, password))))

	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", rancherToken))

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("error resetting password")
	}

	// send email
	err = utils.SendEmail(
		creometryGmail,
		email,
		// needs to be changed to the actual creometry gmail password
		gmailPassword,
		"Creometry Password Reset",
		"Your password has been reset successfully.",
	)

	if err != nil {
		return err
	}

	return nil
}

// Local functions

func createRancherProject(usrProjectName string, plan string) (string, int64, string, error) {
	nsDefaultResourceQuotaLimit, resourceQuotaLimit := genLimitsFromPlan(plan)
	if nsDefaultResourceQuotaLimit == nil && resourceQuotaLimit == nil {
		return "", 0, "", fmt.Errorf("invalid plan")
	}
	resourceQuota := genResourceQuota(*nsDefaultResourceQuotaLimit, *resourceQuotaLimit)

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

func genResourceQuota(nsDefaultResourceQuotaLimit Limit, resourceQuota Limit) string {

	return fmt.Sprintf(`
	"namespaceDefaultResourceQuota": {
		"limit": {
		"configMaps": "%d",
		"limitsCpu": "%s",
		"limitsMemory": "%s",
		"persistentVolumeClaims": "%d",
		"pods": "%d",
		"replicationControllers": "%d",
		"requestsStorage": "%s",
		"secrets": "%d",
		"services": "%d",
		"servicesLoadBalancers": "%d",
		"servicesNodePorts": "%d"
		}
		},
		"resourceQuota": {
		"limit": {
		"configMaps": "%d",
		"limitsCpu": "%s",
		"limitsMemory": "%s",
		"persistentVolumeClaims": "%d",
		"pods": "%d",
		"replicationControllers": "%d",
		"requestsStorage": "%s",
		"secrets": "%d",
		"services": "%d",
		"servicesLoadBalancers": "%d",
		"servicesNodePorts": "%d"
		},
		"usedLimit": { }
		}
	`, nsDefaultResourceQuotaLimit.ConfigMaps, nsDefaultResourceQuotaLimit.LimitsCpu, nsDefaultResourceQuotaLimit.LimitsMemory, nsDefaultResourceQuotaLimit.PersistentVolumeClaims, nsDefaultResourceQuotaLimit.Pods, nsDefaultResourceQuotaLimit.ReplicationControllers, nsDefaultResourceQuotaLimit.RequestsStorage, nsDefaultResourceQuotaLimit.Secrets, nsDefaultResourceQuotaLimit.Services, nsDefaultResourceQuotaLimit.ServicesLoadBalancers, nsDefaultResourceQuotaLimit.ServicesNodePorts, resourceQuota.ConfigMaps, resourceQuota.LimitsCpu, resourceQuota.LimitsMemory, resourceQuota.PersistentVolumeClaims, resourceQuota.Pods, resourceQuota.ReplicationControllers, resourceQuota.RequestsStorage, resourceQuota.Secrets, resourceQuota.Services, resourceQuota.ServicesLoadBalancers, resourceQuota.ServicesNodePorts)
}

func genLimitsFromPlan(plan string) (*Limit, *Limit) {
	switch plan {
	case STARTER:
		return &Limit{ConfigMaps: 10, LimitsCpu: "1000m", LimitsMemory: "2000Mi", PersistentVolumeClaims: 10, Pods: 50, ReplicationControllers: 15, RequestsStorage: "50000Mi", Secrets: 20, Services: 50, ServicesLoadBalancers: 0, ServicesNodePorts: 0}, &Limit{ConfigMaps: 10, LimitsCpu: "1000m", LimitsMemory: "2000Mi", PersistentVolumeClaims: 10, Pods: 100, ReplicationControllers: 30, RequestsStorage: "50000Mi", Secrets: 20, Services: 50, ServicesLoadBalancers: 0, ServicesNodePorts: 0}
	case PRO:
		return &Limit{ConfigMaps: 20, LimitsCpu: "2000m", LimitsMemory: "4000Mi", PersistentVolumeClaims: 20, Pods: 100, ReplicationControllers: 25, RequestsStorage: "50000Mi", Secrets: 20, Services: 50, ServicesLoadBalancers: 0, ServicesNodePorts: 0}, &Limit{ConfigMaps: 20, LimitsCpu: "2000m", LimitsMemory: "4000Mi", PersistentVolumeClaims: 20, Pods: 100, ReplicationControllers: 25, RequestsStorage: "50000Mi", Secrets: 20, Services: 50, ServicesLoadBalancers: 0, ServicesNodePorts: 0}

	case ELITE:
		return &Limit{ConfigMaps: 20, LimitsCpu: "4000m", LimitsMemory: "8000Mi", PersistentVolumeClaims: 30, Pods: 200, ReplicationControllers: 50, RequestsStorage: "200000Mi", Secrets: 20, Services: 100, ServicesLoadBalancers: 0, ServicesNodePorts: 0}, &Limit{ConfigMaps: 20, LimitsCpu: "4000m", LimitsMemory: "8000Mi", PersistentVolumeClaims: 30, Pods: 200, ReplicationControllers: 50, RequestsStorage: "200000Mi", Secrets: 20, Services: 100, ServicesLoadBalancers: 0, ServicesNodePorts: 0}
	}
	return nil, nil
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
		Balance: req.Balance,
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

func createKeyCloakUser(username, email, firstName, lastName, realm, password string) error {
	ctx := context.Background()

	user := &keycloak.User{
		Enabled:       keycloak.Bool(true),
		Username:      keycloak.String(username),
		Email:         keycloak.String(email),
		FirstName:     keycloak.String(firstName),
		LastName:      keycloak.String(lastName),
		EmailVerified: keycloak.Bool(false),
		RequiredActions: []string{
			"UPDATE_PASSWORD",
		},
	}

	res, err := auth.K.Users.Create(ctx, realm, user)
	if err != nil {
		return err
	}

	if res.StatusCode != 201 {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	location := res.Header.Get("location")
	// extract the user id from the location header, it is the last part of the url
	userId := location[len(location)-36:]

	res, err = auth.K.Users.ResetPassword(ctx, realm, userId, &keycloak.Credential{
		Type:      keycloak.String("password"),
		Value:     keycloak.String(password),
		Temporary: keycloak.Bool(true),
	})

	if err != nil {
		return err
	}

	if res.StatusCode != 204 {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	log.Printf("user %s created\nID: %s", username, userId)

	return nil
}

/*func registerKeycloakUser(code string) (string, error) {

	rancherURL, err := utils.GetVariable("config", "RANCHER_URL")
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v3-public/keyCloakProviders/keycloak?action=login", rancherURL), bytes.NewBuffer([]byte(`{"code":"`+code+`"}`)))

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
	dt := ExchangeCodeToTokenResponse{}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &dt)

	if err != nil {
		return "", err
	}

	return dt.AccessToken, nil

}*/

type region struct {
	Name    string `json:"name"`
	Cluster string `json:"cluster"`
}

type project struct {
	Name string `json:"name"`
}

var regions = []region{
	{
		Name:    "us-east-1",
		Cluster: "us-east-1",
	},
}

var projects = []project{
	{Name: "default"},
}

func getUserProjectsByRegion(region, clusterId, userId string) []project {
	// returns all the projects in a region that the user has access to
	return projects
}
func getAllUserProjects(username string) []project {
	// for each region, get the userId and get all the projects that the user has access to
	prs := []project{}
	for _, region := range regions {
		//1
		userId, _, _ := GetUserByUsername(username, region.Name)
		//2
		p := getUserProjectsByRegion(region.Name, region.Cluster, userId)
		//3
		prs = append(prs, p...)
	}
	//4
	return prs
}
