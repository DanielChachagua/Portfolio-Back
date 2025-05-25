package controllers

import (
	"os"
	"path/filepath"

	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/gofiber/fiber/v2"
)

func GetImageProject(c *fiber.Ctx) error {
	name := c.Params("name")
	if name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Image name is required",
		})
	}

	// Construct the file path
	filePath := filepath.Join("./images", name)
	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Image not found",
		})
	}
	// Serve the file
	return c.SendFile(filePath, false)
}