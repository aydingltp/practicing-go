package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"practicing-go/database"
	"practicing-go/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Static("/", "./client/dist")
	database.ConnectDb()
	router.SetupRoutes(app)

	app.Listen(":3000")
}
