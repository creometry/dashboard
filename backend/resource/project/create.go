package project

import (
	"bytes"
	"context"
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

func CreateProject(usrProjectName string, plan string) (kubeconfig string, err error) {

	projectId, err := createRancherProject(usrProjectName)
	if err != nil {
		return "", err
	}

	// create a new namespace with annotation "projectId"
	nsClient := auth.MyClientSet.CoreV1().Namespaces()

	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf(usrProjectName),
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
	switch plan {
	case "silver":
		quota = createResourseQuota("1", "1Gi", newNs.Name)
	case "gold":
		quota = createResourseQuota("2", "2Gi", newNs.Name)
	case "platinum":
		quota = createResourseQuota("4", "4Gi", newNs.Name)
	}

	newQuota, err := quotaClient.Create(context.TODO(), quota, metav1.CreateOptions{})

	if err != nil {
		return "", err
	}
	log.Println("Created Quota:", newQuota.Name)


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
