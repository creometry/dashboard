package namespace

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Creometry/dashboard/auth"
	v1 "k8s.io/api/core/v1"
)


func GetNamespaces() []v1.Namespace {
	ns, err := auth.MyClientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		log.Fatal(err)
	}
	return ns.Items
}
