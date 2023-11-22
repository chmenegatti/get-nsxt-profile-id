package main

import (
	"log"

	"github.com/chmenegatti/get-nsxt-profile-id/database"
	"github.com/chmenegatti/get-nsxt-profile-id/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	routes.Routes(app)

	log.Fatal(app.Listen(":3333"))
}
