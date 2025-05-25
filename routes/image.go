package routes

import (
	"github.com/DanielChachagua/Portfolio-Back/controllers"
	"github.com/gofiber/fiber/v2"
)

func ImageRoutes(app *fiber.App){
	image := app.Group("/image")
	image.Get("/get/:name", controllers.GetImageProject)
}