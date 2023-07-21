package controllers

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/db"
	m "github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	// Declare a variable to hold the request body
	var data map[string]string

	// Parse the request body into the data variable
	err := c.BodyParser(&data)
	if err != nil {
		// If there's an error in parsing the body, log the error and return an Internal Server Error response
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}

	// Declare a variable to hold the user data
	var user *m.User

	result := db.DB.Where(&m.User{Email: data["email"]}).Find(&user)
	if user.IsDeleted {
		return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Your Account is Deleted"})
	}
	// Check if a user with the same email exists
	// If a user with the same email already exists, return a response indicating that
	if result.RowsAffected == 1 {
		return c.Status(fiber.StatusOK).JSON(&m.Response{Success: false, Message: "E-Mail already exists"})
	}

	// Hash the password provided by the user
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)

	// Create a new user with the data provided in the request body
	user = &m.User{
		ID:       uuid.New(),
		Name:     data["name"],
		Email:    data["email"],
		Password: string(password),
	}

	// Save the new user to the database
	tx := db.DB.Create(&user)
	if tx.Error != nil {
		// If there's an error in saving the user, log the error and return an Internal Server Error response
		log.Println(tx.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Error Creating Account"})
	}

	// If the user was saved successfully, return a success response
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "User Created Successfully"})
}

func LoginUser(c *fiber.Ctx) error {
	// Load .env file to get the environment variables
	err := godotenv.Load()
	if err != nil {
		// Log an error message if the .env file could not be loaded
		log.Println("Error loading .env files.")
		log.Fatal(err.Error())
	}

	// Declare a variable to hold the request body
	var data map[string]string

	// Parse the request body into the data variable
	err = c.BodyParser(&data)
	if err != nil {
		// If there's an error in parsing the body, log the error and return an Internal Server Error response
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}

	// Declare a variable to hold the user data
	var user m.User

	// Check if a user with the same email exists
	result := db.DB.Where(&m.User{Email: data["email"]}).First(&user)
	// If the email does not exist in the database, return a response indicating that
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusOK).JSON(&m.Response{Success: false, Message: "E-Mail Not Found"})
	}

	if user.IsDeleted {
		return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Your Account is Deleted"})
	}

	// Compare the hashed password in the database with the password provided by the user
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"]))
	if err != nil {
		// If the passwords do not match, log the error and return a response indicating that
		log.Println(err.Error())
		return c.Status(fiber.StatusOK).JSON(&m.Response{Success: false, Message: "Incorrect Password"})
	}

	// Checking if user's account is active or not
	if user.IsDeactivated {
		return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Account is not active. Please activate your account."})
	}

	// Generate a JWT token for the user, adding the user's ID as a claim and setting an expiration time
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
	})
	// Sign the token with the secret key
	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		// If there's an error in signing the token, log the error
		log.Println(err.Error())
	}
	logs := &m.AccountStatusLogs{
		ID:     uuid.New(),
		UserID: user.ID,
		Status: "loggedin",
	}
	result = db.DB.Create(&logs)
	if result.RowsAffected == 0 {
		log.Println(result.Error)
	}

	// If the login was successful, return a success response along with the JWT token
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Login Successful", Data: fiber.Map{"token": token}})
}

func GetUser(c *fiber.Ctx) error {
	// Get the user_id from the local context and cast it to a string
	user_id := c.Locals("user_id").(string)

	// Parse the user_id into a UUID
	id, err := uuid.Parse(user_id)

	// If error occurs parsing the used_id return Internal Server Error
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}

	// Define a user variable of type m.User
	var user m.User

	// Query the database for a user with the given ID. First(&user) will order by primary key and limit the result to the first record.
	// The user data is then loaded into the 'user' object.
	result := db.DB.Where(&m.User{ID: id}).First(&user)

	// If no rows are affected by the query (i.e., the user was not found in the database), then return a 404 status and a JSON response indicating that the user was not found.
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&m.Response{Success: false, Message: "User Not Found"})
	}

	if user.IsDeleted {
		return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Your Account is Deleted"})
	}
	logs := &m.AccountStatusLogs{
		ID:     uuid.New(),
		UserID: id,
		Status: "fetched",
	}
	result = db.DB.Create(&logs)
	if result.RowsAffected == 0 {
		log.Println(result.Error)
	}

	// If the user was found, return a 200 status and a JSON response indicating that the user was found, along with the user data.
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "User Found Successfully", Data: user})

}

func DeactivateUser(c *fiber.Ctx) error {
	// Get the user_id from the local context and cast it to a string
	user_id := c.Locals("user_id").(string)

	// Parse the user_id into a UUID
	id, err := uuid.Parse(user_id)

	// If error occurs parsing the used_id return Internal Server Error
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}

	// Define a user variable of type m.User
	var user m.User

	// Query the database for a user with the given ID. First(&user) will order by primary key and limit the result to the first record.
	// The user data is then loaded into the 'user' object.
	result := db.DB.Where(&m.User{ID: id}).First(&user)

	// If no rows are affected by the query (i.e., the user was not found in the database), then return a 404 status and a JSON response indicating that the user was not found.
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&m.Response{Success: false, Message: "User Not Found"})
	}

	if user.IsDeleted {
		return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Your Account is Deleted"})
	}

	if user.IsDeactivated {
		return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Account Already Deactivated"})
	}
	now := time.Now()
	user.IsDeactivated = true
	user.DeactivatedAt = &now

	result = db.DB.Save(&user)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusOK).JSON(&m.Response{Success: false, Message: "Error Deactivating Account"})
	}
	logs := &m.AccountStatusLogs{
		ID:     uuid.New(),
		UserID: id,
		Status: "deactivated",
	}
	result = db.DB.Create(&logs)
	if result.RowsAffected == 0 {
		log.Println(result.Error)
	}
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Account Deactivated Successfully"})
}

func ActivateUser(c *fiber.Ctx) error {
	// Get the user_id from the local context and cast it to a string
	user_id := c.Locals("user_id").(string)

	// Parse the user_id into a UUID
	id, err := uuid.Parse(user_id)

	// If error occurs parsing the used_id return Internal Server Error
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}

	// Define a user variable of type m.User
	var user m.User

	// Query the database for a user with the given ID. First(&user) will order by primary key and limit the result to the first record.
	// The user data is then loaded into the 'user' object.
	result := db.DB.Where(&m.User{ID: id}).First(&user)

	// If no rows are affected by the query (i.e., the user was not found in the database), then return a 404 status and a JSON response indicating that the user was not found.
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&m.Response{Success: false, Message: "User Not Found"})
	}

	if user.IsDeleted {
		return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Your Account is Deleted"})
	}

	if !user.IsDeactivated {
		return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Account Already Activated"})
	}
	user.IsDeactivated = false
	user.DeactivatedAt = nil

	result = db.DB.Save(&user)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusOK).JSON(&m.Response{Success: false, Message: "Error Activating Account"})
	}
	logs := &m.AccountStatusLogs{
		ID:     uuid.New(),
		UserID: id,
		Status: "activated",
	}
	result = db.DB.Create(&logs)
	if result.RowsAffected == 0 {
		log.Println(result.Error)
	}
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Account Activated Successfully"})
}

func DeleteUser(c *fiber.Ctx) error {
	// Get the user_id from the local context and cast it to a string
	user_id := c.Locals("user_id").(string)

	// Parse the user_id into a UUID
	id, err := uuid.Parse(user_id)

	// If error occurs parsing the used_id return Internal Server Error
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}

	// Define a user variable of type m.User
	var user m.User

	// Query the database for a user with the given ID. First(&user) will order by primary key and limit the result to the first record.
	// The user data is then loaded into the 'user' object.
	result := db.DB.Where(&m.User{ID: id}).First(&user)

	// If no rows are affected by the query (i.e., the user was not found in the database), then return a 404 status and a JSON response indicating that the user was not found.
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&m.Response{Success: false, Message: "User Not Found"})
	}

	if user.IsDeleted {
		return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Account Already Deleted"})
	}
	now := time.Now()
	user.IsDeleted = true
	user.DeletedAt = &now

	result = db.DB.Save(&user)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusOK).JSON(&m.Response{Success: false, Message: "Error Deleting Account"})
	}
	logs := &m.AccountStatusLogs{
		ID:     uuid.New(),
		UserID: id,
		Status: "deleted",
	}
	result = db.DB.Create(&logs)
	if result.RowsAffected == 0 {
		log.Println(result.Error)
	}
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Account Deleted Successfully"})
}

func BulkCreateUser(c *fiber.Ctx) error {
	file, err := c.FormFile("users")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&m.Response{Success: false, Message: "Error Uploading File"})
	}
	src, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Failed to open the uploaded file"})
	}
	defer src.Close()

	// Create a CSV reader
	reader := csv.NewReader(src)
	// Process the CSV data
	var users []*m.User
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Failed to read CSV file"})
		}
		if record[0] == "name" && record[1] == "email" && record[2] == "password" && record[3] == "role_id" {
			continue
		} else {
			name := record[0]
			email := record[1]
			password := record[2]
			role_id, _ := uuid.Parse(record[3])
			tx := db.DB.Where(&m.User{Email: email}).Find(&m.User{})
			if tx.RowsAffected == 0 {
				hashed_password, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

				// Create a new user with the data provided in the request body
				user := &m.User{
					ID:       uuid.New(),
					Name:     name,
					Email:    email,
					Password: string(hashed_password),
					RoleID:   role_id,
				}

				users = append(users, user)
			}
			// Process each record here
			fmt.Println(record)
		}
	}
	// Save the new user to the database
	tx := db.DB.Create(&users)
	if tx.Error != nil {
		// If there's an error in saving the user, log the error and return an Internal Server Error response
		log.Println(tx.Error)
	}
	fmt.Println(users)
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Users Created Successfully"})
}
