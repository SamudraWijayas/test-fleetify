package routes

import (
	"procurement/config"
	"procurement/handlers"
	"procurement/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	app.Use(middleware.Protected())

	app.Get("/items", handlers.GetItems(config.DB))
	app.Post("/items", handlers.CreateItem(config.DB))

	app.Get("/suppliers", handlers.GetSuppliers(config.DB))
	app.Post("/suppliers", handlers.CreateSupplier(config.DB))

	app.Post("/purchases", handlers.CreatePurchase)
}
