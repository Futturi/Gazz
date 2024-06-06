package handler

import (
	"github.com/Futturi/Gaz/internal/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"log/slog"
	"strings"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body entities.User true "account info"
// @Success 200 {string} string "id"
// @Failure 400 {string} string "incorrect data"
// @Failure 500 {string} string "smth wrong in server"
// @Router /auth/signup [post]
func (h *Handler) SignUp(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		slog.Error("error with parsing json", "error", err)
		return c.Status(fasthttp.StatusBadRequest).JSON(fiber.Map{
			"error": "incorrect data",
		})
	}
	if user.Username == "" || user.Email == "" || user.Password == "" || user.Birthday == "" {
		slog.Error("error with parsing json")
		return c.Status(fasthttp.StatusBadRequest).JSON(fiber.Map{
			"error": "incorrect data",
		})
	}
	if !strings.Contains(user.Email, "@") {
		slog.Error("inccorrect email")
		return c.Status(fasthttp.StatusBadRequest).JSON(fiber.Map{
			"error": "incorrect mail",
		})
	}
	id, err := h.service.SignUp(user)
	if err != nil {
		slog.Error("smth wrong", "error", err)
		return c.Status(fasthttp.StatusInternalServerError).JSON(fiber.Map{
			"error": "smth wrong in server",
		})
	}
	slog.Info("user was added", "id", id)
	return c.Status(fasthttp.StatusOK).JSON(fiber.Map{
		"id": id,
	})
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body entities.User true "account info"
// @Success 200 {string} string "token"
// @Failure 400 {string} string "incorrect data"
// @Failure 500 {string} string "smth wrong in server"
// @Router /auth/signin [post]
func (h *Handler) SignIn(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		slog.Error("error with parsing json", "error", err)
		return c.Status(fasthttp.StatusBadRequest).JSON(fiber.Map{
			"error": "incorrect data",
		})
	}
	if user.Password == "" || user.Email == "" {
		slog.Error("incorrect request")
		return c.Status(fasthttp.StatusBadRequest).JSON(fiber.Map{
			"error": "incorrect data",
		})
	}
	token, err := h.service.SignIn(user)
	if err != nil {
		slog.Error("error with logging", "error", err)
		return c.Status(fasthttp.StatusInternalServerError).JSON(fiber.Map{
			"error": "smth wrong in server",
		})
	}
	slog.Info("user was logged with", "token", token)
	return c.Status(fasthttp.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
