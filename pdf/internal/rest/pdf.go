package rest

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type ResponseError struct {
	Message string `json:"message"`
}

type PdfService interface {
	CompressPDFService(tempPath, outPath string) error
}

type PdfHandler struct {
	Service PdfService
}

func NewPdfHandler(e *echo.Echo, svc PdfService) {
	handler := &PdfHandler{
		Service: svc,
	}

	e.POST("/pdf/compress", handler.CompressPDF)
}

func (p *PdfHandler) CompressPDF(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Invalid file"})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: "Failed to open file"})
	}
	defer src.Close()

	tempPath := filepath.Join("./temp", file.Filename)
	outPath := filepath.Join("./temp", "compressed_"+file.Filename)

	dst, err := os.Create(tempPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: "Failed to create temp file"})
	}
	defer dst.Close()

	if _, err := dst.ReadFrom(src); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: "Failed to save file"})
	}

	if err := p.Service.CompressPDFService(tempPath, outPath); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: "Compression failed"})
	}

	return c.Attachment(outPath, "compressed_"+file.Filename)
}
