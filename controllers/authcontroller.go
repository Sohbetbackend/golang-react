package controllers

import(
	"github.com/gofiber/fiber/v2"
	"github.com/Sohbetbackend/golang-react/database"
	"github.com/Sohbetbackend/golang-react/models"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name: data["name"],
		Email: data["email"],
		Password: password,
	}

	database.DB.Create(&user)

	return c.JSON(data)
}
