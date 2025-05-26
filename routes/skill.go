package routes

import (
	"github.com/DanielChachagua/Portfolio-Back/controllers"
	"github.com/gofiber/fiber/v2"
)

func SkillRoutes(app *fiber.App){
	skill := app.Group("/skill")
	skill.Get("/getAll", controllers.SkillGetAll)
	skill.Post("/create", controllers.SkillCreate)
	skill.Put("/update/:id", controllers.SkillUpdate)
	skill.Delete("/delete/:id", controllers.SkillDelete)
}