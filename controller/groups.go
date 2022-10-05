package controller

import (
	"strings"

	"github.com/chmenegatti/get-nsxt-profile-id/iresponse"
	"github.com/chmenegatti/get-nsxt-profile-id/service"
	"github.com/gofiber/fiber/v2"
)

func GroupsId(c *fiber.Ctx) error {
	var sufix iresponse.Sufix

	edge := c.Params("edge")

	if err := c.BodyParser(&sufix); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	resp, err := service.GetGroupsIds(edge)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var switchIds []service.Group

	for r := range resp.Results {
		if strings.Contains(resp.Results[r].DisplayName, sufix.SufixDev) ||
			strings.Contains(resp.Results[r].DisplayName, sufix.SufixProd) {
			switchIds = append(switchIds, resp.Results[r])
		}
	}

	return c.Status(200).JSON(switchIds)
}
