package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/luthfi/inventory-pos-app/backend/repositories"
)

func AuthRequired(repo repositories.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(401).JSON(fiber.Map{"error": "Format Authorization harus 'Bearer <token>'"})
		}
		tokenString := parts[1]

		secretKey := []byte(os.Getenv("JWT_SECRET"))
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid Token"})
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["id"].(float64))
		c.Locals("userId", userID)

		if c.Path() == "/logout" {
			return c.Next()
		}

		user, _ := repo.FindByID(userID)
		if user.ActiveToken != tokenString {
			return c.Status(401).JSON(fiber.Map{"error": "Sesi telah berakhir"})
		}

		c.Locals("userId", userID)

		return c.Next()
	}
}
