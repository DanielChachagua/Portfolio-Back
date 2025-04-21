package middleware

import (
	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/DanielChachagua/Portfolio-Back/services"
	"github.com/DanielChachagua/Portfolio-Back/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Obtener el token del encabezado de autorización
		token := c.Get("Authorization")

		// Verificar si el token es válido
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token no proporcionado",
			})
		}

		claims, err := utils.VerifyToken(token)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token inválido",
			})
		}

		userId := claims.(jwt.MapClaims)["id"].(string)

		user, err := services.CurrentUser(userId)

		if err != nil {
			if errResp, ok := err.(*models.ErrorStruc); ok {
				return c.Status(errResp.StatusCode).JSON(models.Response{
					Status:  false,
					Body:    nil,
					Message: errResp.Message,
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: "Error interno",
			})
		}
		
		c.Locals("user", user)

		return c.Next()
	}
}