package service

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Creometry/dashboard/auth"
	v1 "k8s.io/api/core/v1"
)


func  GetServices(namespace string) ([]v1.Service, error) {
	services, err := auth.MyClientSet.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	return services.Items, err
}

func  GetService(namespace string, serviceName string) (v1.Service, error) {
	service, err := auth.MyClientSet.CoreV1().Services(namespace).Get(context.TODO(), serviceName, metav1.GetOptions{})
	return *service, err
}
