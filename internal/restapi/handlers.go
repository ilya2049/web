package restapi

import (
	"github.com/gofiber/fiber/v2"

	"web/internal/pkg/stdconv"
	"web/internal/product"
	"web/internal/user"
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
	productIDString := c.Params(id)
	productID := stdconv.ParseInt(productIDString)

	product := h.service.GetProductByID(productID)

	return c.JSON(product)
}

func (h ProductHandler) updateProduct(c *fiber.Ctx) error {
	productIDString := c.Params(id)
	productID := stdconv.ParseInt(productIDString)

	product := product.Product{}
	if err := c.BodyParser(&product); err != nil {
		return err
	}

	h.service.UpdateProduct(productID, product)

	return nil
}

func (h ProductHandler) deleteProduct(c *fiber.Ctx) error {
	productIDString := c.Params(id)
	productID := stdconv.ParseInt(productIDString)

	h.service.DeleteProduct(productID)

	return nil
}

func NewUserHandler(service *user.Service) UserHandler {
	return UserHandler{
		service: service,
	}
}

type UserHandler struct {
	service *user.Service
}

func (h UserHandler) addUser(c *fiber.Ctx) error {
	user := user.User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	h.service.AddUser(user)

	return nil
}

func (h UserHandler) getUser(c *fiber.Ctx) error {
	users := h.service.GetUsers()

	return c.JSON(users)
}

func (h UserHandler) getUserByID(c *fiber.Ctx) error {
	userIDString := c.Params(id)
	userID := stdconv.ParseInt(userIDString)

	user := h.service.GetUsersByID(userID)

	return c.JSON(user)
}

func (h UserHandler) updateUser(c *fiber.Ctx) error {
	userIDString := c.Params(id)
	userID := stdconv.ParseInt(userIDString)

	user := user.User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	h.service.UpdateUser(userID, user)

	return nil
}

func (h UserHandler) deleteUser(c *fiber.Ctx) error {
	userIDString := c.Params(id)
	userID := stdconv.ParseInt(userIDString)

	h.service.DeleteUser(userID)

	return nil
}
