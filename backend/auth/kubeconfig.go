package auth

import (
	"flag"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var MyClientSet *kubernetes.Clientset

func CreateKubernetesClient(){
	kubeconfig := flag.String("kubeconfig", "./kubeconfig.yaml", "")
	flag.Parse()
	config,err:=clientcmd.BuildConfigFromFlags("",*kubeconfig)	

	if err != nil {
		log.Fatal(err)
	}

	clientset,err:=kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	MyClientSet=clientset
}