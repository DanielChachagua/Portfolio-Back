package routes

import (
	"github.com/DanielChachagua/Portfolio-Back/controllers"
	"github.com/DanielChachagua/Portfolio-Back/middleware"
	"github.com/gofiber/fiber/v2"
)

func SkillRoutes(app *fiber.App){
	skill := app.Group("/skill")
	skill.Get("/getAll", controllers.SkillGetAll)
	skill.Post("/create", middleware.AuthMiddleware(), controllers.SkillCreate)
	skill.Put("/update/:id", middleware.AuthMiddleware(), controllers.SkillUpdate)
	skill.Delete("/delete/:id", middleware.AuthMiddleware(), controllers.SkillDelete)
}