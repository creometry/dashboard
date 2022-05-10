package persistentvolume

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Creometry/dashboard/auth"
	v1 "k8s.io/api/core/v1"
)


func GetPersistentVolumes(namespace string) ([]v1.PersistentVolume, error) {

	pvClient := auth.MyClientSet.CoreV1().PersistentVolumes()

	list, err := pvClient.List(context.TODO(), metav1.ListOptions{})
	return list.Items, err

}

func GetPersistentVolume(namespace string, pvName string) (v1.PersistentVolume, error) {

	pvClient := auth.MyClientSet.CoreV1().PersistentVolumes()

	pv, err := pvClient.Get(context.TODO(), pvName, metav1.GetOptions{})
	return *pv, err

}