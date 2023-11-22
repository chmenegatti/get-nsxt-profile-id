package controller

import (
	"github.com/chmenegatti/get-nsxt-profile-id/database"
	"github.com/chmenegatti/get-nsxt-profile-id/iresponse"
	"github.com/chmenegatti/get-nsxt-profile-id/models"
	"github.com/gofiber/fiber/v2"
)

func CreateOrganization(c *fiber.Ctx) error {
	var organization models.Organizations

	if err := c.BodyParser(&organization); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Save(&organization)

	responseOrganization := iresponse.CreateResponse(organization)

	return c.Status(200).JSON(responseOrganization)
}
