package controllers

import (
	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/DanielChachagua/Portfolio-Back/services"
	"github.com/gofiber/fiber/v2"
)


func GetUser (c *fiber.Ctx) error {
	baseUrl := c.BaseURL()
	user, err := services.GetUser(baseUrl)
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
	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  true,
		Body:    user,
		Message: "User retrieved successfully",
	})
}