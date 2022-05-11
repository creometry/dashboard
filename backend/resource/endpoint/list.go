package endpoint

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Creometry/dashboard/auth"
	v1 "k8s.io/api/core/v1"
)

func GetEndpoints(namespace string) ([]v1.Endpoints, error) {

	endpointsClient := auth.MyClientSet.CoreV1().Endpoints(namespace)

	list, err := endpointsClient.List(context.TODO(), metav1.ListOptions{})
	return list.Items, err

}

func GetEndpoint(namespace string, endpointName string) (v1.Endpoints, error) {

	endpointsClient := auth.MyClientSet.CoreV1().Endpoints(namespace)

	endpoint, err := endpointsClient.Get(context.TODO(), endpointName, metav1.GetOptions{})
	return *endpoint, err

}
