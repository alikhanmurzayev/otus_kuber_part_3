package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func StartServer(port int, userController *userController) error {
	app := fiber.New()

	app.Get("/health", userController.HealthCheck)
	userGroup := app.Group("/user")
	userGroup.Post("/", userController.CreateUser)
	userGroup.Get("/:id", userController.GetUser)
	userGroup.Delete("/:id", userController.DeleteUser)
	userGroup.Put("/:id", userController.UpdateUser)

	return app.Listen(":" + strconv.Itoa(port))
}
