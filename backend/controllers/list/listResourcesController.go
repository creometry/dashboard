package controllers

import (
	"github.com/Creometry/dashboard/resource/configmap"
	"github.com/Creometry/dashboard/resource/cronjob"
	"github.com/Creometry/dashboard/resource/customresource"
	"github.com/Creometry/dashboard/resource/deployment"
	"github.com/Creometry/dashboard/resource/endpoint"
	"github.com/Creometry/dashboard/resource/event"
	"github.com/Creometry/dashboard/resource/horizontalpodautoscaler"
	"github.com/Creometry/dashboard/resource/ingress"
	"github.com/Creometry/dashboard/resource/job"
	"github.com/Creometry/dashboard/resource/persistentvolumeclaim"
	"github.com/Creometry/dashboard/resource/pod"
	"github.com/Creometry/dashboard/resource/project"
	"github.com/Creometry/dashboard/resource/secret"
	"github.com/Creometry/dashboard/resource/service"
	"github.com/Creometry/dashboard/resource/statefulset"

	"github.com/gofiber/fiber/v2"
)


func GetAllPods(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	pods, err := pod.GetPods(ns)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": pods,
	})
}

func GetPod(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	podName := c.Params("pod")
	pod, err := pod.GetPod(ns, podName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": pod,
	})
}

func GetAllServices(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	services, err := service.GetServices(ns)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": services,
	})
}

func GetService(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	serviceName := c.Params("service")
	service, err := service.GetService(ns, serviceName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": service,
	})
}

func GetAllDeployments(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	deployments, err := deployment.GetDeployments(ns)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": deployments,
	})
}

func GetDeployment(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	deploymentName := c.Params("deployment")
	deployment, err := deployment.GetDeployment(ns, deploymentName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": deployment,
	})
}

func GetAllConfigMaps(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	configMaps, err := configmap.GetConfigMaps(ns)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": configMaps,
	})
}

func GetConfigMap(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	configMapName := c.Params("configmap")
	configMap, err := configmap.GetConfigMap(ns, configMapName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": configMap,
	})
}

func GetAllSecrets(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	secrets, err := secret.GetSecrets(ns)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": secrets,
	})
}

func GetSecret(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	secretName := c.Params("secret")
	secret, err := secret.GetSecret(ns, secretName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": secret,
	})
}

func GetAllPersistentVolumeClaims(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	pvc, err := persistentvolumeclaim.GetPersistentVolumeClaims(ns)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": pvc,
	})
}

func GetPersistentVolumeClaim(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	pvcName := c.Params("pvc")
	pvc, err := persistentvolumeclaim.GetPersistentVolumeClaim(ns, pvcName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": pvc,
	})
}

func GetAllStatefulSets(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	statefulSets, err := statefulset.GetStatefulSets(ns)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": statefulSets,
	})
}

func GetStatefulSet(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	statefulSetName := c.Params("sts")
	statefulSet, err := statefulset.GetStatefulSet(ns, statefulSetName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": statefulSet,
	})
}

func GetAllJobs(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	jobs, err := job.GetJobs(ns)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": jobs,
	})
}

func GetJob(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	jobName := c.Params("job")
	job, err := job.GetJob(ns, jobName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": job,
	})
}

func GetAllCronJobs(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	cronJobs, err := cronjob.GetCronJobs(ns)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": cronJobs,
	})
}

func GetCronJob(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	cronJobName := c.Params("cronjob")
	cronJob, err := cronjob.GetCronJob(ns, cronJobName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": cronJob,
	})
}

func GetAllEndpoints(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	endpoints, err := endpoint.GetEndpoints(ns)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": endpoints,
	})
}

func GetEndpoint(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	endpointName := c.Params("endpoint")
	endpoint, err := endpoint.GetEndpoint(ns, endpointName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": endpoint,
	})
}

func GetAllIngresses(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	ingresses, err := ingress.GetIngresses(ns)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": ingresses,
	})
}

func GetIngress(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	ingressName := c.Params("ingress")
	ingress, err := ingress.GetIngress(ns, ingressName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": ingress,
	})
}

func GetAllEvents(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	events, err := event.GetEvents(ns)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": events,
	})
}

func GetAllHorizontalPodAutoscalers(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	horizontalPodAutoscalers, err := horizontalpodautoscaler.GetHorizontalPodAutoscalers(ns)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": horizontalPodAutoscalers,
	})
}

func GetHorizontalPodAutoscaler(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	horizontalPodAutoscalerName := c.Params("horizontalpodautoscaler")
	horizontalPodAutoscaler, err := horizontalpodautoscaler.GetHorizontalPodAutoscaler(ns, horizontalPodAutoscalerName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": horizontalPodAutoscaler,
	})
}

func GetAllCustomResources(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	customResources, err := customresource.GetCustomResources(ns)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": customResources,
	})
}

func GetCustomResource(c *fiber.Ctx) error {
	ns := c.Params("namespace")
	customResourceName := c.Params("customresource")
	customResource, err := customresource.GetCustomResource(ns, customResourceName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": customResource,
	})
}


func GetNampespacesByAnnotation(c *fiber.Ctx) error {
	annotation := c.Params("annotation")
	namespace,projectId, err := project.GetNamespaceByAnnotation([]string{annotation})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"namespace": namespace,
		"projectId": projectId,
	})
}