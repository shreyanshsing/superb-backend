package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"superb.com/v0/user_service/models"
)

var Database *gorm.DB

func ConnectToDB() {

	err := godotenv.Load("../.env") // Adjust the path as necessary
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
	// Get the database connection details from environment variables
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Connect to the PostgreSQL database
	_db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Database = _db

	_db.AutoMigrate(&models.UserModel{})

	fmt.Println("Connected to the PostgreSQL database!")
}
