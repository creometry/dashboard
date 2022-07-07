package persistentvolumeclaim

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Creometry/resources-service/auth"
	v1 "k8s.io/api/core/v1"
)

func GetPersistentVolumeClaims(namespace string) ([]v1.PersistentVolumeClaim, error) {

	pvcClient := auth.MyClientSet.CoreV1().PersistentVolumeClaims(namespace)

	list, err := pvcClient.List(context.TODO(), metav1.ListOptions{})
	return list.Items, err

}

func GetPersistentVolumeClaim(namespace string, pvcName string) (v1.PersistentVolumeClaim, error) {

	pvcClient := auth.MyClientSet.CoreV1().PersistentVolumeClaims(namespace)

	pvc, err := pvcClient.Get(context.TODO(), pvcName, metav1.GetOptions{})
	return *pvc, err

}
