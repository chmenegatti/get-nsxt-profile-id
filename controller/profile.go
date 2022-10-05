package controller

import (
	"github.com/chmenegatti/get-nsxt-profile-id/service"
	"github.com/gofiber/fiber/v2"
)

type ProfileId struct {
	ID      string `json:"id,omitempty"`
	Segment string `json:"segment"`
	Message string `json:"message,omitempty"`
}

func GetProfileId(c *fiber.Ctx) error {
	var profile []ProfileId

	edge := c.Params("edge")

	if err := c.BodyParser(&profile); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var segment []ProfileId

	for i := 0; i < len(profile); i++ {

		response, err := service.GetProfileIds(profile[i].Segment, edge)

		if len(response.Results) != 0 {
			if err != nil {
				return c.Status(400).JSON(err.Error())
			}
			profile[i].ID = response.Results[0].Id

			segment = append(segment, profile[i])
		} else {
			profile[i].Message = "Não existe no NSXT"
			segment = append(segment, profile[i])
		}

	}

	return c.Status(200).JSON(segment)

}
