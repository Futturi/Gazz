package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"log/slog"
)

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	id := c.Locals("user").(jwt.MapClaims)["id"].(float64)
	if id == 0 {
		slog.Error("unauthorized")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	users, err := h.service.GetUsers(id)
	if err != nil {
		slog.Error("smth wrong", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "smth wrong in server",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"users": users,
	})
}
