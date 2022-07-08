package customresource

import (
	"context"
	"encoding/json"

	"github.com/Creometry/resources-service/auth"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetCustomResources(namespace string) ([]interface{}, error) {
	crdClient := auth.MyExtensionsClientSet.ApiextensionsV1().CustomResourceDefinitions()

	list, err := crdClient.List(context.TODO(), metav1.ListOptions{})
	var resList []interface{}
	var element interface{}

	for _, crd := range list.Items {

		restClient, err := NewRESTClient(auth.Config, &crd)
		if err != nil {
			return nil, err
		}

		raw, err := restClient.Get().NamespaceIfScoped(namespace, crd.Spec.Scope == apiextensionsv1.NamespaceScoped).Resource(crd.Spec.Names.Plural).Do(context.TODO()).Raw()

		if err != nil {
			return nil, err
		}

		// unmarshal the raw json into the list of resources (inumplemented)
		err = json.Unmarshal(raw, &element)

		if err != nil {
			return nil, err
		}

		// add the custom resources to the list
		resList = append(resList, element)

	}

	// return the list of custom resources
	return resList, err
}

func GetCustomResource(namespace string, crdName string) (interface{}, error) {
	crdClient := auth.MyExtensionsClientSet.ApiextensionsV1().CustomResourceDefinitions()

	crd, err := crdClient.Get(context.TODO(), crdName, metav1.GetOptions{})
	var element interface{}

	restClient, err := NewRESTClient(auth.Config, crd)
	if err != nil {
		return nil, err
	}

	raw, err := restClient.Get().NamespaceIfScoped(namespace, crd.Spec.Scope == apiextensionsv1.NamespaceScoped).Resource(crd.Spec.Names.Plural).Do(context.TODO()).Raw()

	if err != nil {
		return nil, err
	}

	// unmarshal the raw json into the list of resources (inumplemented)
	err = json.Unmarshal(raw, &element)

	if err != nil {
		return nil, err
	}

	return element, err
}
