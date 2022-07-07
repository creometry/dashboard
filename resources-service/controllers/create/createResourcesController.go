package create

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"strings"
	"time"

	"github.com/Creometry/resources-service/auth"
	"github.com/gofiber/fiber/v2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateNamespace(c *fiber.Ctx) error {

	// get the request body
	var req CreateNsRequestBody
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// validate the request body
	if err := req.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// create a new namespace with annotation "projectId"
	nsClient := auth.MyInClusterClientSet.CoreV1().Namespaces()

	// create a random hash and append it to the namespace name
	h := sha1.New()
	h.Write([]byte(time.Now().String()))
	b := h.Sum(nil)
	rand := base64.URLEncoding.EncodeToString(b)
	// replace every special character in the random hash with a random letter
	rand = strings.Replace(rand, "+", "x", -1)
	rand = strings.Replace(rand, "/", "x", -1)
	rand = strings.Replace(rand, "=", "x", -1)
	rand = strings.Replace(rand, ".", "x", -1)
	rand = strings.Replace(rand, "_", "x", -1)
	rand = strings.Replace(rand, "*", "x", -1)
	rand = strings.Replace(rand, " ", "x", -1)
	rand = strings.Replace(rand, ",", "x", -1)

	nsName := strings.ToLower(req.ProjectName) + "-" + strings.ToLower(rand)

	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: nsName,
			Annotations: map[string]string{
				"field.cattle.io/projectId": req.ProjectId,
			},
			Labels: map[string]string{
				"field.cattle.io/projectId": strings.Split(req.ProjectId, ":")[1],
			},
		},
	}

	newNs, err := nsClient.Create(context.TODO(), ns, metav1.CreateOptions{})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"ns_name": newNs.Name,
	})
}
