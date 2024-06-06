package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"log/slog"
	"strings"
)

const (
	salt2 = "legjliqjwoejrqfgeniowo4i3wipreq;ksdfjbgkhoiterwpq[lasdkmcvnbfjghutriowepq[als;dk,m,cvnkfjg"
)

func JwtMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing or malformed JWT",
		})
	}

	token, err := jwt.Parse(strings.Split(tokenString, " ")[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		return []byte(salt2), nil
	})

	if err != nil || !token.Valid {
		slog.Error("error with token", "error", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid or expired JWT",
		})
	}
	c.Locals("user", token.Claims.(jwt.MapClaims))

	return c.Next()
}
