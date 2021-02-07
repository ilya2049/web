package restapi

import (
	"web/internal/product"
	"web/internal/user"

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

	userService := user.NewService()
	userHandler := NewUserHandler(&userService)

	app.Post("/users", userHandler.addUser)
	app.Get("/users", userHandler.getUser)
	app.Get("/users/:id", userHandler.getUserByID)
	app.Put("/users/:id", userHandler.updateUser)
	app.Delete("/users/:id", userHandler.deleteUser)

	app.Listen(":3000")
}
