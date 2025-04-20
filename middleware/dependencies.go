package middleware

import (
	"context"

	"github.com/DanielChachagua/Portfolio-Back/dependencies"
	ctxt "github.com/DanielChachagua/Portfolio-Back/context"
	"github.com/gofiber/fiber/v2"
)

func InjectApp(app *dependencies.Dependency) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext() 
		ctx = context.WithValue(ctx, ctxt.AppKey, app) 
		c.SetUserContext(ctx)

		return c.Next()
	}
}