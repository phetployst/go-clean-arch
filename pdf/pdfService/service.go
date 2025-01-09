package pdfService

import (
	"fmt"

	pdfApi "github.com/pdfcpu/pdfcpu/pkg/api"
)

type PdfRepository interface{}

type Service struct {
	pdfRepository PdfRepository
}

func NewService(r PdfRepository) *Service {
	return &Service{
		pdfRepository: r,
	}
}

func (s *Service) CompressPDFService(tempPath, outPath string) error {
	err := pdfApi.OptimizeFile(tempPath, outPath, nil)
	if err != nil {
		return fmt.Errorf("failed to compress PDF: %w", err) // wrap error เพื่อช่วยในการ debug
	}
	return nil
}
