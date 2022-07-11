package pod

import (
	"context"

	"github.com/Creometry/resources-service/auth"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPods(namespace string) ([]v1.Pod, error) {

	podsClient := auth.MyInClusterClientSet.CoreV1().Pods(namespace)
	pods, err := podsClient.List(context.TODO(), metav1.ListOptions{})
	return pods.Items, err

}

func GetPod(namespace string, podName string) (v1.Pod, error) {

	podsClient := auth.MyInClusterClientSet.CoreV1().Pods(namespace)
	pod, err := podsClient.Get(context.TODO(), podName, metav1.GetOptions{})
	return *pod, err

}
