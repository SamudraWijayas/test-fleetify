package handlers

import (
	"time"

	"procurement/config"
	"procurement/models"

	"github.com/gofiber/fiber/v2"
)

type PurchaseRequest struct {
	SupplierID uint
	Items      []struct {
		ItemID uint
		Qty    int
	}
}

func CreatePurchase(c *fiber.Ctx) error {
	var req PurchaseRequest
	c.BodyParser(&req)

	tx := config.DB.Begin()

	purchase := models.Purchasing{
		Date:       time.Now(),
		SupplierID: req.SupplierID,
		UserID:     1,
	}

	if err := tx.Create(&purchase).Error; err != nil {
		tx.Rollback()
		return err
	}

	grandTotal := 0

	for _, i := range req.Items {
		var item models.Item
		tx.First(&item, i.ItemID)

		if item.Stock < i.Qty {
			tx.Rollback()
			return c.Status(400).JSON(fiber.Map{"error": "stock not enough"})
		}

		subTotal := item.Price * i.Qty
		grandTotal += subTotal

		tx.Create(&models.PurchasingDetail{
			PurchasingID: purchase.ID,
			ItemID:       item.ID,
			Qty:          i.Qty,
			SubTotal:     subTotal,
		})

		tx.Model(&item).Update("stock", item.Stock-i.Qty)
	}

	tx.Model(&purchase).Update("grand_total", grandTotal)
	tx.Commit()

	return c.JSON(fiber.Map{"message": "purchase success"})
}
