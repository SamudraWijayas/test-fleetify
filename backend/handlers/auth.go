package handlers

import (
	"time"

	"procurement/config"
	"procurement/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var body models.User
	c.BodyParser(&body)

	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	body.Password = string(hash)

	config.DB.Create(&body)
	return c.JSON(fiber.Map{"message": "register success"})
}

func Login(c *fiber.Ctx) error {
	var body models.User
	var user models.User

	c.BodyParser(&body)
	config.DB.Where("username = ?", body.Username).First(&user)

	if user.ID == 0 {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	t, _ := token.SignedString(config.JwtSecret())

	return c.JSON(fiber.Map{"token": t})
}
