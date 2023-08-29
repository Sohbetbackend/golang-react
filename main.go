package main

import (
	"github.com/Sohbetbackend/golang-react/database"
	"github.com/Sohbetbackend/golang-react/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/middleware/cors"
)

func main() {
	database.Connect()

    app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

    app.Listen(":8000")
}