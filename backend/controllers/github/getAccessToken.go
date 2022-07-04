package controllers

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func GetAccessToken(c *fiber.Ctx) error {
	ctx := context.Background()
	// get request params
	code := c.Params("code")
	// check if the request params are valid
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "code is required",
		})
	}
	conf := &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes:       []string{"public_repo", "user"},
		Endpoint:     github.Endpoint,
	}
	// get access token
	token, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Println("errr")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error exchanging code for token",
		})
	}
	return c.JSON(fiber.Map{
		"access_token": token.AccessToken,
	})
}
