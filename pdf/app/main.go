package main

import (
	"log"

	"github.com/labstack/echo/v4"
	rest "github.com/phetployst/go-clean-arch/pdf/internal/rest"
	pdfService "github.com/phetployst/go-clean-arch/pdf/pdfService"
)

const defaultAddress = ":8080"

func main() {
	e := echo.New()

	service := pdfService.NewService()
	rest.NewPdfHandler(e, service)

	if err := e.Start(defaultAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
