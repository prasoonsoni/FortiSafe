package middlewares

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/db"
	m "github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func AuthenticateUser(c *fiber.Ctx) error {
	// Get the Authorization header from the HTTP request.
	authHeader := c.Get("Authorization")

	// If the Authorization header is empty, return a 401 Unauthorized status.
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(&m.Response{Success: false, Message: "No Authorization header provided"})
	}

	// Split the Authorization header into two parts based on the "Bearer " string.
	headerParts := strings.Split(authHeader, "Bearer ")

	// If the Authorization header does not have two parts (i.e., "Bearer " and the JWT), return a 401 Unauthorized status.
	if len(headerParts) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(&m.Response{Success: false, Message: "Invalid Authorization header format. Format is 'Bearer <token>'"})
	}

	// Extract the JWT from the second part of the Authorization header.
	jwtToken := headerParts[1]

	// Parse the JWT.
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Check if the token's signing method is what we expect.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the JWT secret for HMAC verification.
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	// If there was an error in parsing the JWT (which could mean the JWT is invalid or expired), return a 401 Unauthorized status.
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(&m.Response{Success: false, Message: "Invalid or expired JWT"})
	}

	// If the JWT's claims can be successfully converted into MapClaims and the token is valid, add the user_id from the claims to the request context and call the next handler in the chain.
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user_id, _ := uuid.Parse(claims["user"].(string))
		var user *m.User
		_ = db.DB.Where(&m.User{ID: user_id}).Find(&user)
		c.Locals("user_role", user.RoleID)
		c.Locals("user_id", claims["user_id"])
		return c.Next()
	}

	// If the JWT's claims could not be converted into MapClaims or the token was invalid, return a 401 Unauthorized status.
	return c.Status(fiber.StatusUnauthorized).JSON(&m.Response{Success: false, Message: "Invalid or expired JWT"})

}
