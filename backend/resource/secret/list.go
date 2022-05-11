package secret

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Creometry/dashboard/auth"
	v1 "k8s.io/api/core/v1"
)

func GetSecrets(namespace string) ([]v1.Secret, error) {

	secretsClient := auth.MyClientSet.CoreV1().Secrets(namespace)

	list, err := secretsClient.List(context.TODO(), metav1.ListOptions{})
	return list.Items, err

}

func GetSecret(namespace string, secretName string) (v1.Secret, error) {

	secretsClient := auth.MyClientSet.CoreV1().Secrets(namespace)

	secret, err := secretsClient.Get(context.TODO(), secretName, metav1.GetOptions{})
	return *secret, err

}
