package gormRepo

import (
	"github.com/phetployst/go-clean-arch/pdf/pdfService"
	"gorm.io/gorm"
)

type PdfRepository struct {
	DB *gorm.DB
}

func NewPdfRepository(db *gorm.DB) pdfService.PdfRepository {
	return &PdfRepository{
		DB: db,
	}
}
