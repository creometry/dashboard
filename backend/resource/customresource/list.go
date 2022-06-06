package customresource

import (
	"context"
	"runtime"

	"github.com/Creometry/dashboard/auth"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func GetCustomResources(namespace string)([]interface{}, error) {
	crdClient := auth.MyExtensionsClientSet.ApiextensionsV1().CustomResourceDefinitions()

	list, err := crdClient.List(context.TODO(),metav1.ListOptions{})
	var res runtime.Object
	var resList []interface{}

	for _,crd := range list.Items {
		crd = []apiextensionsv1.CustomResourceDefinition{removeNonServedVersions(crd)}[0]
		restClient, err := NewRESTClient(auth.Config, &crd)
		if err != nil {
			return nil, err
		}

		err= restClient.Get().NamespaceIfScoped(namespace,crd.Spec.Scope == apiextensionsv1.NamespaceScoped).Resource(crd.Spec.Names.Plural).Do(context.TODO()).Into(&res)

		if err != nil {
			return nil, err
		}

		resList = append(res, res)
		
	}

	
	return resList, err
}

func removeNonServedVersions(crd apiextensionsv1.CustomResourceDefinition) apiextensionsv1.CustomResourceDefinition {
	versions := make([]apiextensions.CustomResourceDefinitionVersion, 0)

	for _, version := range crd.Spec.Versions {
		if version.Served {
			versions = append(versions, version)
		}
	}

	crd.Spec.Versions = versions
	return crd
}