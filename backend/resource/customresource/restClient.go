package customresource

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
)

func getCustomResourceDefinitionGroupVersion(crd *apiextensions.CustomResourceDefinition) schema.GroupVersion {
	if len(crd.Spec.Versions) == 0 {
		return schema.GroupVersion{
			Group:   crd.Spec.Group,
			Version: "v1",
		}
	}
	return schema.GroupVersion{
		Group:   crd.Spec.Group,
		Version: crd.Spec.Versions[0].Name,
	}
}


func NewRESTClient(config *rest.Config, crd *apiextensions.CustomResourceDefinition) (*rest.RESTClient, error) {
	groupVersion := getCustomResourceDefinitionGroupVersion(crd)
	scheme := runtime.NewScheme()
	schemeBuilder := runtime.NewSchemeBuilder(
		func(scheme *runtime.Scheme) error {
			scheme.AddKnownTypes(
				groupVersion,
				&metav1.ListOptions{},
				&metav1.DeleteOptions{},
			)
			return nil
		})
	if err := schemeBuilder.AddToScheme(scheme); err != nil {
		return nil, err
	}

	config.GroupVersion = &groupVersion
	config.APIPath = "/apis"
	config.ContentType = runtime.ContentTypeJSON
	config.NegotiatedSerializer = serializer.WithoutConversionCodecFactory{CodecFactory: serializer.NewCodecFactory(scheme)}

	return rest.RESTClientFor(config)
}