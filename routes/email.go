package routes

import (
	"github.com/DanielChachagua/Portfolio-Back/controllers"
	"github.com/gofiber/fiber/v2"
)

func EmailRoutes(app *fiber.App) {
	email := app.Group("/email")
	email.Post("/send_email", controllers.SendEmail)
}