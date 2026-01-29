package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"procurement/models"
)

func GetSuppliers(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var suppliers []models.Supplier
		db.Find(&suppliers)
		return c.JSON(suppliers)
	}
}

func CreateSupplier(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var supplier models.Supplier
		c.BodyParser(&supplier)
		db.Create(&supplier)
		return c.JSON(supplier)
	}
}
