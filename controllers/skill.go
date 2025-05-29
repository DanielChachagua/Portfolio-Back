package controllers

import (
	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/DanielChachagua/Portfolio-Back/services"
	"github.com/gofiber/fiber/v2"
)

func SkillCreate(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Failed to parse image file",
		})
	}

	if file.Size > 1*1024*1024 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: "La imagen no debe superar 1MB",
		})
	}

	var skillCreate models.SkillCreate
	skillCreate.Name = c.FormValue("name")
	skillCreate.Area = c.FormValue("area")

	if err := skillCreate.Validate(); err != nil {
		return c.Status(422).JSON(models.Response{
			Status: false,
			Message: "Invalid model",
			Body: nil,
		})
	}

	id, err := services.SkillCreate(&skillCreate, file)

	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		return c.Status(500).JSON(models.Response{
			Status: false,
			Message: "Error creating skill",
			Body: nil,
		})
	}

	return c.Status(201).JSON(models.Response{
		Status:  true,
		Body:    id,
		Message: "Skill created successfully",
	})
}

func SkillGetAll(c *fiber.Ctx) error {
	baseUrl := c.BaseURL()
	skills, err := services.SkillGetAll(baseUrl)
	if err != nil {
		return c.Status(500).JSON(models.Response{
			Status: false,
			Message: "Error getting skills",
			Body: nil,
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    skills,
		Message: "Skills retrieved successfully",
	})
}

func SkillUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(models.Response{
			Status: false,
			Message: "ID is required",
			Body: nil,
		})
	}

	file, _ := c.FormFile("image")

	if file.Size > 1*1024*1024 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: "La imagen no debe superar 1MB",
		})
	}

	var skillUpdate models.SkillUpdate
	skillUpdate.Name = c.FormValue("name")
	skillUpdate.Area = c.FormValue("area")

	if err := skillUpdate.Validate(); err != nil {
		return c.Status(422).JSON(models.Response{
			Status: false,
			Message: "Invalid model",
			Body: nil,
		})
	}

	err := services.SkillUpdate(id, &skillUpdate, file)
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		return c.Status(500).JSON(models.Response{
			Status: false,
			Message: "Error updating skill",
			Body: nil,
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Message: "Skill updated successfully",
	})
}

func SkillDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(models.Response{
			Status: false,
			Message: "ID is required",
			Body: nil,
		})
	}

	err := services.SkillDelete(id)
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		return c.Status(500).JSON(models.Response{
			Status: false,
			Message: "Error deleting skill",
			Body: nil,
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Message: "Skill deleted successfully",
	})
}