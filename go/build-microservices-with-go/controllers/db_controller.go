package controllers

import (
	"microservice_example/db/orm"
	"microservice_example/db/plain_sql"

	"github.com/gofiber/fiber/v2"
)

func GetNameSQL(c *fiber.Ctx) error {
	/* Param parsing and processing */
	id := c.Params("id")

	/* Search in DB */
	name, err := plain_sql.SQLFindByNConst(id)
	if err != nil {
		return err
	}

	/* Successful JSON response */
	return c.JSON(fiber.Map{
		"success": true,
		"name":    name,
	})
}

func GetNameSQLX(c *fiber.Ctx) error {
	/* Param parsing and processing */
	id := c.Params("id")

	/* Search in DB */
	name, err := plain_sql.SQLXFindByNConst(id)
	if err != nil {
		return err
	}

	/* Successful JSON response */
	return c.JSON(fiber.Map{
		"success": true,
		"name":    name,
	})
}

func GetNamePGX(c *fiber.Ctx) error {
	/* Param parsing and processing */
	id := c.Params("id")

	/* Search in DB */
	name, err := plain_sql.PGXFindByNConst(id)
	if err != nil {
		return err
	}

	/* Successful JSON response */
	return c.JSON(fiber.Map{
		"success": true,
		"name":    name,
	})
}

func GetNameGorm(c *fiber.Ctx) error {
	/* Param parsing and processing */
	id := c.Params("id")

	/* Search in DB */
	name, err := orm.GormFindByNConst(id)
	if err != nil {
		return err
	}

	/* Successful JSON response */
	return c.JSON(fiber.Map{
		"success": true,
		"name":    name,
	})
}
