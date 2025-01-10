package rest

import (
	"bytes"
	"io"
	"mime/multipart"
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

		mockService.On("CompressPDFService", mock.Anything, mock.Anything).Return(nil)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		file, err := writer.CreateFormFile("file", "test.pdf")
		assert.NoError(t, err)

		_, err = io.Copy(file, bytes.NewReader([]byte("mock pdf content")))
		assert.NoError(t, err)

		err = writer.Close()
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/pdf/compress", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		HandlerErr := handler.CompressPDF(c)

		assert.NoError(t, HandlerErr)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
