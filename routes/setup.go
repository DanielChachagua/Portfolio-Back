package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	AuthRoutes(app)
	ProjectRoutes(app)
	SkillRoutes(app)
	UserRoutes(app)
	ImageRoutes(app)
	EmailRoutes(app)
}