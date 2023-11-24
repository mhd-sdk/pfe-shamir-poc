package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pfe-manager/config"
	"github.com/pfe-manager/pkg/servicesManager"
	"github.com/pfe-manager/pkg/shamir"
)

type saveSecretBody struct {
	Secret string `json:"secret"`
}

func HandleSaveSecret(c *fiber.Ctx) error {
	secret := new(saveSecretBody)
	if err := c.BodyParser(secret); err != nil {
		return c.Status(400).SendString("Bad Request")
	}
	split, err := shamir.SplitSeed([]byte(secret.Secret), config.GetConfig().ShamirSplitNumber, config.GetConfig().ShamirThreshold)
	if err != nil {
		return c.Status(500).SendString("Internal Server Error: Could not split secret")
	}

	err = servicesManager.Dispatch(split)
	if err != nil {
		return c.Status(500).SendString("Internal Server Error: Could not dispatch secret: " + err.Error())
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Secret saved",
	})
}

func HandleGetSecret(c *fiber.Ctx) error {
	c.Status(200)
	secret, err := servicesManager.Retrieve()
	if err != nil {
		return c.Status(500).SendString("Internal Server Error: Could not retrieve secret: " + err.Error())
	}
	return c.JSON(fiber.Map{
		"secret": string(secret),
	})
}
