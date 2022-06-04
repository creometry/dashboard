package cronjob

import (
	"context"

	"github.com/Creometry/dashboard/auth"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetCronJobs(namespace string) ([]batchv1.CronJob, error) {

	cronJobsClient := auth.MyClientSet.BatchV1().CronJobs(namespace)

	list, err := cronJobsClient.List(context.TODO(), metav1.ListOptions{})
	return list.Items, err

}

func GetCronJob(namespace string, cronJobName string) (batchv1.CronJob, error) {

	cronJobsClient := auth.MyClientSet.BatchV1().CronJobs(namespace)

	cronJob, err := cronJobsClient.Get(context.TODO(), cronJobName, metav1.GetOptions{})
	return *cronJob, err

}
