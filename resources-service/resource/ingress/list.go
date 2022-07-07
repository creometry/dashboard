package ingress

import (
	"context"

	"github.com/Creometry/resources-service/auth"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetIngresses(namespace string) (v1.IngressList, error) {

	ingressesClient := auth.MyClientSet.NetworkingV1().Ingresses(namespace)

	list, err := ingressesClient.List(context.TODO(), metav1.ListOptions{})
	return *list, err

}

func GetIngress(namespace string, ingressName string) (v1.Ingress, error) {

	ingressesClient := auth.MyClientSet.NetworkingV1().Ingresses(namespace)

	ingress, err := ingressesClient.Get(context.TODO(), ingressName, metav1.GetOptions{})
	return *ingress, err

}
