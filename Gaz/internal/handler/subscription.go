package handler

import (
	"github.com/Futturi/Gaz/internal/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"log/slog"
)

// @Summary Subscribe
// @Security ApiKeyAuth
// @Tags subscription
// @Description subscribe
// @ID subscribe
// @Accept  json
// @Produce  json
// @Param request body entities.SubscribeReq true "request"
// @Success 200 {string} string "subscribed"
// @Failure 400 {string} string "incorrect data"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "smth wrong in server"
// @Router /api/subscribe [post]
func (h *Handler) Subscribe(c *fiber.Ctx) error {
	id := c.Locals("user").(jwt.MapClaims)["id"].(float64)
	if id == 0 {
		slog.Error("unauthorized")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	var req entities.SubscribeReq

	if err := c.BodyParser(&req); err != nil {
		slog.Error("error with parsing json", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "incorrect data",
		})
	}

	err := h.service.Subscribe(id, req.Username)
	if err != nil {
		slog.Error("smth wrong", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "smth wrong in server",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "subscribed",
	})
}

// @Summary Unsubscribe
// @Security ApiKeyAuth
// @Tags unsubscription
// @Description unsubscribe
// @ID unsubscribe
// @Accept  json
// @Produce  json
// @Param request body entities.SubscribeReq true "request"
// @Success 200 {string} string "unsubscribed"
// @Failure 400 {string} string "incorrect data"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "smth wrong in server"
// @Router /api/unsubscribe [post]
func (h *Handler) Unsubscribe(c *fiber.Ctx) error {
	id := c.Locals("user").(jwt.MapClaims)["id"].(float64)
	if id == 0 {
		slog.Error("unauthorized")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	var req entities.SubscribeReq

	if err := c.BodyParser(&req); err != nil {
		slog.Error("error with parsing json", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "incorrect data",
		})
	}

	err := h.service.Unsubscribe(id, req.Username)
	if err != nil {
		slog.Error("smth wrong", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "smth wrong in server",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "unsubscribed",
	})
}
