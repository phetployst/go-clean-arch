package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	gormRepo "github.com/phetployst/go-clean-arch/pdf/internal/repository/gorm"
	rest "github.com/phetployst/go-clean-arch/pdf/internal/rest"
	pdfService "github.com/phetployst/go-clean-arch/pdf/pdfService"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultAddress = ":8080"

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	dbConnStr := os.Getenv("DB_CONNECTION_STRING")
	if dbConnStr == "" {
		log.Fatal("DB_CONNECTION_STRING is not set in the environment variables")
	}

	db, err := gorm.Open(postgres.Open(dbConnStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	repo := gormRepo.NewPdfRepository(db)
	service := pdfService.NewService(repo)
	rest.NewPdfHandler(e, service)

	if err := e.Start(defaultAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
