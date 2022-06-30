package horizontalpodautoscaler

import (
	"context"

	"github.com/Creometry/dashboard/auth"
	autoscaling "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetHorizontalPodAutoscalers(namespace string) ([]autoscaling.HorizontalPodAutoscaler, error) {

	horizontalPodAutoscalersClient := auth.MyClientSet.AutoscalingV1().HorizontalPodAutoscalers(namespace)

	list, err := horizontalPodAutoscalersClient.List(context.TODO(), metav1.ListOptions{})
	return list.Items, err
}

func GetHorizontalPodAutoscaler(namespace string, name string) (autoscaling.HorizontalPodAutoscaler, error) {

	horizontalPodAutoscalersClient := auth.MyClientSet.AutoscalingV1().HorizontalPodAutoscalers(namespace)

	hpo, err := horizontalPodAutoscalersClient.Get(context.TODO(), name, metav1.GetOptions{})
	return *hpo, err
}
