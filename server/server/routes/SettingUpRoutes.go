package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/root/medicalBooking/controllers"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/user", controllers.GetAllUsers)
	app.Get("/user/:id", controllers.GetSingleUser)
	app.Post("/user", controllers.CreateUser)
	app.Put("/user/:id", controllers.UpdateUser)
	app.Delete("/user/:id", controllers.DeleteUser)
}
