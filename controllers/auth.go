package controllers

import (
	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/DanielChachagua/Portfolio-Back/services"
	"github.com/gofiber/fiber/v2"
)

func AuthLogin(c *fiber.Ctx) error {
	var loginRequest models.Login

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status: false, 
			Body: nil, 
			Message: "Invalid request",
		})
	}

	if err := loginRequest.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status: false, 
			Body: nil, 
			Message: err.Error(),
		})
	}

	token, err := services.AuthLogin(loginRequest.Email, loginRequest.Password)
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

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    token,
		Message: "Token obtenido con éxito",
	})
}


