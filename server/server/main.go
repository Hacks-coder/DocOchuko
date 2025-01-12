package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/root/medicalBooking/database"
	"github.com/root/medicalBooking/routes"
)

func init() {
	database.ConnectDB()
}
func main() {

	app := fiber.New()
	routes.SetUpRoutes(app)
	app.Listen(":6000")
}
