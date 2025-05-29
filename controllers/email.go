package controllers

import (
	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/DanielChachagua/Portfolio-Back/services"
	"github.com/gofiber/fiber/v2"
)

func SendEmail(c *fiber.Ctx) error {
	var emailContact models.EmailContact
	if err := c.BodyParser(&emailContact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Message: "Invalid request body",
		})
	}
	if err := emailContact.Validate(); err != nil {
		return c.Status(422).JSON(models.Response{
			Status:  false,
			Message: "Validation error",
		})
	}

	if err := services.SendEmail(&emailContact); err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Message: "Error sending email",
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  true,
		Message: "Email sent successfully",
	})
}