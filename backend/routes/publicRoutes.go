package routes

import (
	gh "github.com/Creometry/dashboard/controllers/github"
	controllers "github.com/Creometry/dashboard/controllers/list"
	pr "github.com/Creometry/dashboard/controllers/project"

	"github.com/gofiber/fiber/v2"
)

func CreateRoutes(app *fiber.App) {

	v1 := app.Group("/api/v1")
	v1.Get("/github/exchange/:code",gh.GetAccessToken)
	v1.Post("/project", pr.CreateProject)
	v1.Post("/kubeconfig", pr.GenerateKubeConfig)
	v1.Get("/pods/:namespace", controllers.GetAllPods)
	v1.Get("/pods/:namespace/:pod", controllers.GetPod)
	v1.Get("/services/:namespace", controllers.GetAllServices)
	v1.Get("/services/:namespace/:service", controllers.GetService)
	v1.Get("/deployments/:namespace", controllers.GetAllDeployments)
	v1.Get("/deployments/:namespace/:deployment", controllers.GetDeployment)
	v1.Get("/configmaps/:namespace", controllers.GetAllConfigMaps)
	v1.Get("/configmaps/:namespace/:configmap", controllers.GetConfigMap)
	v1.Get("/secrets/:namespace", controllers.GetAllSecrets)
	v1.Get("/secrets/:namespace/:secret", controllers.GetSecret)
	v1.Get("/pvcs/:namespace", controllers.GetAllPersistentVolumeClaims)
	v1.Get("/pvcs/:namespace/:pvc", controllers.GetPersistentVolumeClaim)
	v1.Get("/sts/:namespace", controllers.GetAllStatefulSets)
	v1.Get("/sts/:namespace/:sts", controllers.GetStatefulSet)
	v1.Get("/jobs/:namespace", controllers.GetAllJobs)
	v1.Get("/jobs/:namespace/:job", controllers.GetJob)
	v1.Get("/cronjobs/:namespace", controllers.GetAllCronJobs)
	v1.Get("/cronjobs/:namespace/:cronjob", controllers.GetCronJob)
	v1.Get("/endpoints/:namespace", controllers.GetAllEndpoints)
	v1.Get("/endpoints/:namespace/:endpoint", controllers.GetEndpoint)
	v1.Get("/ingresses/:namespace", controllers.GetAllIngresses)
	v1.Get("/ingresses/:namespace/:ingress", controllers.GetIngress)
	v1.Get("/events/:namespace", controllers.GetAllEvents)
	v1.Get("/horizontalpodautoscalers/:namespace", controllers.GetAllHorizontalPodAutoscalers)
	v1.Get("/horizontalpodautoscalers/:namespace/:horizontalpodautoscaler", controllers.GetHorizontalPodAutoscaler)
	v1.Get("/customresources/:namespace", controllers.GetAllCustomResources)
	v1.Get("/customresources/:namespace/:customresource", controllers.GetCustomResource)
}
