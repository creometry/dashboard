package configmap

import (
	"context"

	"github.com/Creometry/resources-service/auth"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetConfigMaps(namespace string) ([]v1.ConfigMap, error) {

	configMapsClient := auth.MyInClusterClientSet.CoreV1().ConfigMaps(namespace)

	list, err := configMapsClient.List(context.TODO(), metav1.ListOptions{})
	return list.Items, err

}

func GetConfigMap(namespace string, configMapName string) (v1.ConfigMap, error) {

	configMapsClient := auth.MyInClusterClientSet.CoreV1().ConfigMaps(namespace)

	configMap, err := configMapsClient.Get(context.TODO(), configMapName, metav1.GetOptions{})
	return *configMap, err

}
