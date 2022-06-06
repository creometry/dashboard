package auth

import (
	"flag"
	"log"

	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var MyClientSet *kubernetes.Clientset
var MyExtensionsClientSet *clientset.Clientset
var Config *rest.Config

func CreateKubernetesClient() {
	kubeconfig := flag.String("kubeconfig", "./kubeconfig.yaml", "")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	MyClientSet = clientset
}

func CreateExtensionsClient() {
	config, err := clientcmd.BuildConfigFromFlags("", "./kubeconfig.yaml")

	if err != nil {
		log.Fatal(err)
	}

	clientset,err:=clientset.NewForConfig(config)

	if err != nil {
		log.Fatal(err)
	}

	MyExtensionsClientSet = clientset
	Config = config

}
