package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"practicing-go/controllers"
	"practicing-go/database"
	"practicing-go/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Static("/", "./client/dist")
	database.ConnectDb()

	var db = database.DB()
	controllers.Context(db)
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
