package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"

	"github.com/gofiber/fiber/v2"
)

func Orders(c *fiber.Ctx) error {
	var orders []models.Order

	database.DB.Find(&orders)

	return c.JSON(orders)
}

func CreateOrder(c *fiber.Ctx) error {
}

func DeleteOrder(c *fiber.Ctx) error {
}

func UpdateOrder(c *fiber.Ctx) error {
}
