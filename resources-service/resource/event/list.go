package event

import (
	"context"

	"github.com/Creometry/resources-service/auth"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetEvents(namespace string) ([]v1.Event, error) {

	eventsClient := auth.MyInClusterClientSet.CoreV1().Events(namespace)

	list, err := eventsClient.List(context.TODO(), metav1.ListOptions{})
	return list.Items, err
}
