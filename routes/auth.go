package routes

import (
	"github.com/DanielChachagua/Portfolio-Back/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App){
	auth := app.Group("/auth")
	auth.Post("/login", controllers.AuthLogin)
}