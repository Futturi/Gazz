package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"log/slog"
)

// @Summary Get all users
// @Security ApiKeyAuth
// @Tags users
// @Description get all users
// @ID get-all-users
// @Accept  json
// @Produce  json
// @Success 200 {array} entities.User
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "smth wrong in server"
// @Router /api/users [get]
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
	slog.Info("user with id get all users", "id", id)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"users": users,
	})
}
