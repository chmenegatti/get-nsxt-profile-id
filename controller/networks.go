package controller

import (
	"github.com/chmenegatti/get-nsxt-profile-id/database"
	"github.com/chmenegatti/get-nsxt-profile-id/iresponse"
	"github.com/chmenegatti/get-nsxt-profile-id/models"
	"github.com/gofiber/fiber/v2"
)

func CreateNetwork(c *fiber.Ctx) error {
	var network models.Networks

	if err := c.BodyParser(&network); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&network)

	responseNetwork := iresponse.CreateNetworkResponse(network)

	return c.Status(200).JSON(responseNetwork)
}
