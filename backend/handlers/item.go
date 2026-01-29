package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"procurement/models"
)

func GetItems(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var items []models.Item
		db.Find(&items)
		return c.JSON(items)
	}
}

func CreateItem(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var item models.Item
		c.BodyParser(&item)
		db.Create(&item)
		return c.JSON(item)
	}
}
