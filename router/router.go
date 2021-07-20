package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"practicing-go/controllers"
	"practicing-go/middleware"
)

func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/", Hello)

	login := api.Group("/login")
	login.Get("/", controllers.Login)

	user := api.Group("/user")
	user.Get("/", controllers.GetAllUsers)
	user.Get("/:id", controllers.GetUser)
	user.Post("/", controllers.CreateUser)
	user.Delete("/:id", middleware.Protected(), controllers.DeleteUser)
}
