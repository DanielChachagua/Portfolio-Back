package routes

import (
	"github.com/DanielChachagua/Portfolio-Back/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App){
	user := app.Group("/user")
	user.Get("/get", controllers.GetUser)
}