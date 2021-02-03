package restapi

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func setHeadersMiddleware(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	return c.Next()
}

func loggingMiddleware(c *fiber.Ctx) error {
	log.Println("Path: ", c.Path())

	return c.Next()
}
