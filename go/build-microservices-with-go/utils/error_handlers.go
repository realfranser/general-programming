package utils

import "github.com/gofiber/fiber/v2"

// Returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/404.html")
}
