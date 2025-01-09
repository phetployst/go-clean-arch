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
	// โหลด DB_CONNECTION_STRING และตรวจสอบว่าค่าว่างหรือไม่
	dbConnStr := os.Getenv("DB_CONNECTION_STRING")
	if dbConnStr == "" {
		log.Fatal("DB_CONNECTION_STRING is not set in the environment variables")
	}

	// เปิดการเชื่อมต่อกับฐานข้อมูล
	db, err := gorm.Open(postgres.Open(dbConnStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// ตรวจสอบการทำงานของ AutoMigrate
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	// ตั้งค่า Echo และลงทะเบียน Route
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	// ลงทะเบียน repository, service และ handler
	repo := gormRepo.NewPdfRepository(db)
	service := pdfService.NewService(repo)
	rest.NewPdfHandler(e, service)

	// เริ่มต้นเซิร์ฟเวอร์หลังจากลงทะเบียนทุกอย่างเรียบร้อยแล้ว
	if err := e.Start(defaultAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// ตัวอย่าง struct สำหรับ AutoMigrate
type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"unique"`
}
