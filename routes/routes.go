package routes

import (
	"github.com/chmenegatti/get-nsxt-profile-id/controller"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Post("/:edge/nsxt-profile", controller.GetProfileId)
	app.Post("/organization", controller.CreateOrganization)
	app.Post("/networks", controller.CreateNetwork)
	app.Get("/:edge/switches", controller.SwitchesId)
	app.Get("/:edge/segments", controller.SegmentsIds)
	app.Get("/:edge/groups", controller.GroupsId)
}
