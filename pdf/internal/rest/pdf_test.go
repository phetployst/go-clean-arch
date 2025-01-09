package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPdfService struct {
	mock.Mock
}

func (m *MockPdfService) CompressPDFService(tempPath, outPath string) error {
	args := m.Called(tempPath, outPath)
	return args.Error(0)
}

func TestCompressPDF(t *testing.T) {
	t.Run("compress pdf successfully", func(t *testing.T) {
		mockService := new(MockPdfService)
		handler := &PdfHandler{Service: mockService}

		e := echo.New()
		defer e.Close()

		mockService.On("CompressPDFService", mock.Anything, mock.Anything).Return(nil)

		req := httptest.NewRequest(http.MethodPost, "/pdf/compress", nil)
		req.Form = map[string][]string{
			"file": {"test.pdf"},
		}
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.CompressPDF(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
