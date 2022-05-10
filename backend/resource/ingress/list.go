package ingress

import (
	"context"

	"github.com/Creometry/dashboard/auth"
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetIngresses(namespace string) ([]v1beta1.Ingress, error) {

	ingressesClient := auth.MyClientSet.ExtensionsV1beta1().Ingresses(namespace)

	list, err := ingressesClient.List(context.TODO(), metav1.ListOptions{})
	return list.Items, err

}

func GetIngress(namespace string, ingressName string) (v1beta1.Ingress, error) {

	ingressesClient := auth.MyClientSet.ExtensionsV1beta1().Ingresses(namespace)

	ingress, err := ingressesClient.Get(context.TODO(), ingressName, metav1.GetOptions{})
	return *ingress, err

}