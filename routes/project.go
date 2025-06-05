package routes

import (
	"github.com/DanielChachagua/Portfolio-Back/controllers"
	"github.com/DanielChachagua/Portfolio-Back/middleware"
	"github.com/gofiber/fiber/v2"
)

func ProjectRoutes(app *fiber.App){
	project := app.Group("/project")
	project.Get("/getAll", controllers.GetAllProjects)
	project.Post("/create", middleware.AuthMiddleware(), controllers.CreateProject)
	project.Put("/update/:id", middleware.AuthMiddleware(), controllers.UpdateProject)
	project.Delete("/delete/:id", middleware.AuthMiddleware(), controllers.DeleteProject)
	project.Get("/get/:id", middleware.AuthMiddleware(), controllers.GetProjectByID)
}