package pod

import (
	"context"

	"github.com/Creometry/dashboard/auth"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func  GetPods(namespace string) ([]v1.Pod, error) {
	pods, err := auth.MyClientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	return pods.Items, err
}

func GetPod(namespace string, podName string) (v1.Pod, error) {
	pod, err := auth.MyClientSet.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	return *pod, err
}

