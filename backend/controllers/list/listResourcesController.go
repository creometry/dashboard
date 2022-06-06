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
		"pods": pods,
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
		"pod": pod,
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
		"services": services,
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
		"service": service,
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
		"deployments": deployments,
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
		"deployment": deployment,
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
		"configmaps": configMaps,
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
		"configmap": configMap,
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
		"secrets": secrets,
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
		"secret": secret,
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
		"pvc": pvc,
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
		"pvc": pvc,
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
		"statefulsets": statefulSets,
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
		"statefulset": statefulSet,
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
		"jobs": jobs,
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
		"job": job,
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
		"cronjobs": cronJobs,
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
		"cronjob": cronJob,
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
		"endpoints": endpoints,
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
		"endpoint": endpoint,
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
		"ingresses": ingresses,
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
		"ingress": ingress,
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
		"events": events,
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
		"horizontalpodautoscalers": horizontalPodAutoscalers,
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
		"horizontalpodautoscaler": horizontalPodAutoscaler,
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
		"customresources": customResources,
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
		"customresource": customResource,
	})
}