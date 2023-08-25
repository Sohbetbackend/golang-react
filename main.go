package main

import (
	"github.com/Sohbetbackend/golang-react/database"
	"github.com/Sohbetbackend/golang-react/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
    app := fiber.New()

	routes.Setup(app)

    app.Listen(":3000")
}