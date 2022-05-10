package networkpolicy

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	networking "k8s.io/api/networking/v1"

	"github.com/Creometry/dashboard/auth"
)

func GetNetworkPolicies(namespace string) ([]networking.NetworkPolicy, error) {

	networkPoliciesClient := auth.MyClientSet.NetworkingV1().NetworkPolicies(namespace)
	networkPoliciesm, err := networkPoliciesClient.List(context.TODO(), metav1.ListOptions{})
	return networkPoliciesm.Items, err
}
func GetNetworkPolicy(namespace string,networkPolicyName string ) (networking.NetworkPolicy, error) {
	networkPoliciesClient := auth.MyClientSet.NetworkingV1().NetworkPolicies(namespace)
	networkPolicy, err := networkPoliciesClient.Get(context.TODO(), networkPolicyName, metav1.GetOptions{})
	return *networkPolicy, err
}