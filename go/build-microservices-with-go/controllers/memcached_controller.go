package controllers

import (
	"microservice_example/db/plain_sql"
	"microservice_example/services"

	"github.com/gofiber/fiber/v2"
)

func GetNameMemcache(c *fiber.Ctx) error {
	id := c.Params("id")

	val, err := services.GetName(id)
	if err == nil {
		return c.JSON(fiber.Map{
			"success": true,
			"name":    val,
		})
	}

	name, err := plain_sql.PGXFindByNConst(id)
	if err != nil {
		return err
	}

	_ = services.SetName(name) // XXX

	return c.JSON(fiber.Map{
		"success": true,
		"name":    name,
	})
}
