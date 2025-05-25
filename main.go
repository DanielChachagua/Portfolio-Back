package main

import (
	"log"
	"os"

	"github.com/DanielChachagua/Portfolio-Back/database"
	"github.com/DanielChachagua/Portfolio-Back/dependencies"
	"github.com/DanielChachagua/Portfolio-Back/middleware"
	"github.com/DanielChachagua/Portfolio-Back/repositories"
	"github.com/DanielChachagua/Portfolio-Back/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := database.ConectDB(os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	defer database.CloseDB(db)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: false,
	}))

	dep := dependencies.NewDependency(db)

	app.Use(middleware.LoggingMiddleware)

	routes.SetupRoutes(app)

	repositories.Repo = dep.Repository

	log.Fatal(app.Listen(":4000"))
}
