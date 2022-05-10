package deployment

import (
	"context"

	"github.com/Creometry/dashboard/auth"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)



func GetDeployments(namespace string) ([]appsv1.Deployment, error) {

	deploymentsClient := auth.MyClientSet.AppsV1().Deployments(namespace)

	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	return list.Items, err

}

func  GetDeployment(namespace string, deploymentName string) (appsv1.Deployment, error) {

	deploymentsClient := auth.MyClientSet.AppsV1().Deployments(namespace)

	deployment, err := deploymentsClient.Get(context.TODO(), deploymentName, metav1.GetOptions{})
	return *deployment, err

}