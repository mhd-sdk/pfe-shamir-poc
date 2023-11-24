package api

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pfe-manager/config"
	"github.com/pfe-manager/pkg/api/handlers"
	"github.com/pfe-manager/pkg/api/middlewares"
	"github.com/pfe-manager/pkg/api/websockets"
)

func Init() {
	var disableStartupMessage bool
	if config.GetConfig().Mode == "dev" {
		disableStartupMessage = false
	} else {
		disableStartupMessage = true
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: disableStartupMessage,
	})
	app.Use(cors.New())

	if config.GetConfig().Mode == "dev" {
		app.Use(logger.New())
	}

	app.Use("/ws", middlewares.WebsocketMiddleware)

	app.Post("/api/v1/secret", handlers.HandleSaveSecret)
	app.Get("/api/v1/secret", handlers.HandleGetSecret)

	app.Get("/ws/v1/services/status/:serviceName", websocket.New(websockets.ServiceStatusWS))

	app.Listen("localhost:3000")
}
