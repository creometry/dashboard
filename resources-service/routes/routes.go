package routes

import (
	c "github.com/Creometry/resources-service/controllers/create"
	l "github.com/Creometry/resources-service/controllers/list"

	"github.com/gofiber/fiber/v2"
)

func CreateRoutes(app *fiber.App) {

	v1 := app.Group("/api/v1")
	v1.Post("/namespace", c.CreateNamespace)
	v1.Get("/pods/:namespace", l.GetAllPods)
	v1.Get("/pods/:namespace/:pod", l.GetPod)
	v1.Get("/services/:namespace", l.GetAllServices)
	v1.Get("/services/:namespace/:service", l.GetService)
	v1.Get("/deployments/:namespace", l.GetAllDeployments)
	v1.Get("/deployments/:namespace/:deployment", l.GetDeployment)
	v1.Get("/configmaps/:namespace", l.GetAllConfigMaps)
	v1.Get("/configmaps/:namespace/:configmap", l.GetConfigMap)
	v1.Get("/secrets/:namespace", l.GetAllSecrets)
	v1.Get("/secrets/:namespace/:secret", l.GetSecret)
	v1.Get("/pvcs/:namespace", l.GetAllPersistentVolumeClaims)
	v1.Get("/pvcs/:namespace/:pvc", l.GetPersistentVolumeClaim)
	v1.Get("/sts/:namespace", l.GetAllStatefulSets)
	v1.Get("/sts/:namespace/:sts", l.GetStatefulSet)
	v1.Get("/jobs/:namespace", l.GetAllJobs)
	v1.Get("/jobs/:namespace/:job", l.GetJob)
	v1.Get("/cronjobs/:namespace", l.GetAllCronJobs)
	v1.Get("/cronjobs/:namespace/:cronjob", l.GetCronJob)
	v1.Get("/endpoints/:namespace", l.GetAllEndpoints)
	v1.Get("/endpoints/:namespace/:endpoint", l.GetEndpoint)
	v1.Get("/ingresses/:namespace", l.GetAllIngresses)
	v1.Get("/ingresses/:namespace/:ingress", l.GetIngress)
	v1.Get("/events/:namespace", l.GetAllEvents)
	v1.Get("/horizontalpodautoscalers/:namespace", l.GetAllHorizontalPodAutoscalers)
	v1.Get("/horizontalpodautoscalers/:namespace/:horizontalpodautoscaler", l.GetHorizontalPodAutoscaler)
	v1.Get("/customresources/:namespace", l.GetAllCustomResources)
	v1.Get("/customresources/:namespace/:customresource", l.GetCustomResource)
}
