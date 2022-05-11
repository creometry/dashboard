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


func CreateProject(usrProjectName string) (kubeconfig string, err error) {

	projectId := createRancherProject(usrProjectName)

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

	/*var rs *v1.ResourceQuota

	switch plan {
	case "silver":
		rs = createResourseQuota("1", "1Gi", newNs.Name)
		break
	case "gold":
		rs = createResourseQuota("2", "2Gi", newNs.Name)
		break
	case "platinum":
		rs = createResourseQuota("4", "4Gi", newNs.Name)
		break
	default:
		rs = createResourseQuota("1", "1Gi", newNs.Name)
	}

	// create resource quota for the namespace
	newRs, err := auth.MyClientSet.CoreV1().ResourceQuotas(newNs.Name).Create(context.TODO(), rs, metav1.CreateOptions{})

	if err != nil {
		return "", err
	}
	log.Println("Created ResourceQuota:", newRs.Name)

	// install virtual cluster on the namespace and return kubeconfig*/

	return "kubeconfig", nil

}

func createRancherProject(usrProjectName string) string {
	req, err := http.NewRequest("POST", os.Getenv("CREATE_PROJECT_URL"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"name":"%s","clusterId":"%s"}`, usrProjectName, os.Getenv("CLUSTER_ID")))))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("RANCHER_TOKEN")))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
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
	
	return dt.ProjectId
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
