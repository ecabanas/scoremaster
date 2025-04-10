package database

import (
	"fmt"
	"os"

	"scoremaster/backend/logs"
	"scoremaster/backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func init() {
	godotenv.Load()
}
func InitDatabase() {
	var err error
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbName)
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//DBConn, err = gorm.Open(sqlite.Open("scoremaster.db"), &gorm.Config{})
	if err != nil {
		logs.Error.Println("Failed to connect to database:")
	}
	migrate()
}

func migrate() {
	DBConn.AutoMigrate(
		&models.Category{},
		&models.Question{},
		&models.Answer{},
		&models.Participant{},
		&models.ParticipantAnswer{},
	)
	logs.Info.Println("Database migrated successfully")
}
