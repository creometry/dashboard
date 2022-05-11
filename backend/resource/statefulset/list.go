package statefulset

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Creometry/dashboard/auth"
)

func GetStatefulSets(namespace string) ([]appsv1.StatefulSet, error) {

	statefulSetsClient := auth.MyClientSet.AppsV1().StatefulSets(namespace)

	list, err := statefulSetsClient.List(context.TODO(), metav1.ListOptions{})
	return list.Items, err

}

func GetStatefulSet(namespace string, statefulSetName string) (appsv1.StatefulSet, error) {

	statefulSetsClient := auth.MyClientSet.AppsV1().StatefulSets(namespace)

	statefulSet, err := statefulSetsClient.Get(context.TODO(), statefulSetName, metav1.GetOptions{})
	return *statefulSet, err

}
