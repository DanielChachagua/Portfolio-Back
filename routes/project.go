package routes

import (
	"github.com/DanielChachagua/Portfolio-Back/controllers"
	"github.com/gofiber/fiber/v2"
)

func ProjectRoutes(app *fiber.App){
	project := app.Group("/project")
	project.Get("/getAll", controllers.GetAllProjects)
	project.Post("/create", controllers.CreateProject)
	project.Put("/update/:id", controllers.UpdateProject)
	project.Delete("/delete/:id", controllers.DeleteProject)
	project.Get("/get/:id", controllers.GetProjectByID)
}