package job

import (
	"context"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Creometry/dashboard/auth"
)

func GetJobs(namespace string) ([]batchv1.Job, error) {

	jobsClient := auth.MyClientSet.BatchV1().Jobs(namespace)

	list, err := jobsClient.List(context.TODO(), metav1.ListOptions{})
	return list.Items, err

}

func GetJob(namespace string, jobName string) (batchv1.Job, error) {

	jobsClient := auth.MyClientSet.BatchV1().Jobs(namespace)

	job, err := jobsClient.Get(context.TODO(), jobName, metav1.GetOptions{})
	return *job, err

}