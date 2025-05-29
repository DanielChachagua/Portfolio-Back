package controllers

import (
	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/DanielChachagua/Portfolio-Back/services"
	"github.com/gofiber/fiber/v2"
)

func CreateProject(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Response{
					Status:  false,
					Body:    nil,
					Message: "Failed to parse multipart form",
			})
	}
	// Parse the request body
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

	var project models.CreateProject
	project.Title = c.FormValue("title")
	project.Description = c.FormValue("description")
	project.Link = c.FormValue("link")
	project.Favorite = c.FormValue("favorite") == "true"
	
	if form != nil && form.Value != nil {
			project.SkillsID = form.Value["skills_id"] // []string
	} else {
			project.SkillsID = []string{}
	}

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
	baseUrl := c.BaseURL()
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	project, err := services.GetProjectByID(id, baseUrl)
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
	baseUrl := c.BaseURL()
	projects, err := services.GetAllProjects(baseUrl)
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
	baseUrl := c.BaseURL()
	projects, err := services.GetFavorites(baseUrl)
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
	form, err := c.MultipartForm()
	if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Response{
					Status:  false,
					Body:    nil,
					Message: "Failed to parse multipart form",
			})
	}

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

	if form != nil && form.Value != nil {
			project.SkillsID = form.Value["skills_id"] // []string
	} else {
			project.SkillsID = []string{}
	}

	if err := project.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status: false, 
			Body: nil, 
			Message: err.Error(),
		})
	}

	err = services.UpdateProject(id, file, &project)
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