package restapi

import (
	"github.com/gofiber/fiber/v2"

	"web/internal/pkg/stdconv"
	"web/internal/product"
)

func NewProductHandler(service *product.Service) ProductHandler {
	return ProductHandler{
		service: service,
	}
}

type ProductHandler struct {
	service *product.Service
}

func (h ProductHandler) addProduct(c *fiber.Ctx) error {
	product := product.Product{}
	if err := c.BodyParser(&product); err != nil {
		return err
	}

	h.service.AddProduct(product)

	return nil
}

func (h ProductHandler) getProducts(c *fiber.Ctx) error {
	products := h.service.GetProducts()

	return c.JSON(products)
}

func (h ProductHandler) getProductByID(c *fiber.Ctx) error {
	productIDString := c.Params("id")
	productID := stdconv.ParseInt(productIDString)

	product := h.service.GetProductByID(productID)

	return c.JSON(product)
}
