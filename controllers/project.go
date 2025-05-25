package controllers

import (
	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/DanielChachagua/Portfolio-Back/services"
	"github.com/gofiber/fiber/v2"
)

func CreateProject(c *fiber.Ctx) error {
	// Parse the request body
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Failed to parse image file",
		})
	}

	var project models.CreateProject
	project.Title = c.FormValue("title")
	project.Description = c.FormValue("description")
	project.Link = c.FormValue("link")
	project.Favorite = c.FormValue("favorite") == "true"

	if err := project.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status: false, 
			Body: nil, 
			Message: err.Error(),
		})
	}

	id, err := services.CreateProject(file, &project)
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

	return c.Status(fiber.StatusCreated).JSON(models.Response{
		Status:  true,
		Body:    id,
		Message: "Project created successfully",
	})
}

func GetProjectByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	project, err := services.GetProjectByID(id)
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
		Body:    project,
		Message: "Project retrieved successfully",
	})
}

func GetAllProjects(c *fiber.Ctx) error {
	projects, err := services.GetAllProjects()
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
		Body:    projects,
		Message: "Projects retrieved successfully",
	})
}

func GetFavorites(c *fiber.Ctx) error {
	projects, err := services.GetFavorites()
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
		Body:    projects,
		Message: "Projects retrieved successfully",
	})
}

func UpdateProject(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	file, _ := c.FormFile("image")
	
	var project models.UpdateProject
	project.Title = c.FormValue("title")
	project.Description = c.FormValue("description")
	project.Link = c.FormValue("link")
	project.Favorite = c.FormValue("favorite") == "true"

	if err := project.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status: false, 
			Body: nil, 
			Message: err.Error(),
		})
	}

	err := services.UpdateProject(id, file, &project)
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
		Body:    nil,
		Message: "Project updated successfully",
	})
}

func DeleteProject(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	err := services.DeleteProject(id)
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
		Body:    nil,
		Message: "Project deleted successfully",
	})
}