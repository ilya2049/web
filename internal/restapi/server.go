package restapi

import (
	"web/internal/product"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()

	app.Use(setHeadersMiddleware)
	app.Use(loggingMiddleware)

	productService := product.NewService()
	productHandler := NewProductHandler(&productService)

	app.Post("/products", productHandler.addProduct)
	app.Get("/products", productHandler.getProducts)
	app.Get("/products/:id", productHandler.getProductByID)
	app.Put("/products/:id", productHandler.updateProduct)
	app.Delete("/products/:id", productHandler.deleteProduct)

	app.Listen(":3000")
}
