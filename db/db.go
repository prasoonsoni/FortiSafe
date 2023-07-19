package db

import (
	"fmt"
	"log"
	"os"
	"strconv"

	// "github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Loading environment variables from a .env file
	err := godotenv.Load()
	if err != nil {
		// Log an error message if the .env file could not be loaded
		log.Println("Error loading .env files.")
		log.Fatal(err.Error())
	}

	// Fetching the database host, name, user, and password from the environment variables
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port, _ := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 32)

	// Constructing the data source name (dsn) based on the environment variables
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)

	// Opening a new database connection
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		// If there is an error connecting to the database, logging an error message and terminating the program
		log.Println("Failed to connect to Database")
		log.Fatal(err.Error())
	}

	// err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// When we are successfully connected to the database
	log.Println("Connected to Database Successfully")

}
