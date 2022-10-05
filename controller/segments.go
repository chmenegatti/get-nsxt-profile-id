package controller

import (
	"strings"

	"github.com/chmenegatti/get-nsxt-profile-id/iresponse"
	"github.com/chmenegatti/get-nsxt-profile-id/service"
	"github.com/gofiber/fiber/v2"
)

func SegmentsIds(c *fiber.Ctx) error {
	var sufix iresponse.Sufix

	edge := c.Params("edge")

	if err := c.BodyParser(&sufix); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	resp, err := service.GetSegmentIds(edge)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var segmentIds []service.Segment

	for r := range resp.Results {
		if strings.HasSuffix(resp.Results[r].DisplayName, sufix.SufixDev) ||
			strings.HasSuffix(resp.Results[r].DisplayName, sufix.SufixProd) {
			segmentIds = append(segmentIds, resp.Results[r])
		}
	}

	return c.Status(200).JSON(segmentIds)
}
