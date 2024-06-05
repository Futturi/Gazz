package handler

import (
	"github.com/Futturi/Gaz/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/robfig/cron/v3"
	"log/slog"
)

type Handler struct {
	service *service.Serivce
	cron    *cron.Cron
}

func NewHandler(service *service.Serivce, cr *cron.Cron) *Handler {
	return &Handler{service: service, cron: cr}
}

func (h *Handler) Init() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return true
		},
		AllowCredentials: true,
		MaxAge:           60,
	}))
	_, err := h.cron.AddFunc("30 20 * * *", h.SendMessage)
	if err != nil {
		slog.Error("error with cron", "error", err)
	}
	h.cron.Start()

	auth := app.Group("/auth")

	auth.Post("/signup", h.SignUp)

	auth.Post("/signin", h.SignIn)

	api := app.Group("/api")
	api.Use(jwtMiddleware)
	api.Get("/users", h.GetUsers)

	api.Post("/subscribe", h.Subscribe)

	api.Post("/unsubscribe", h.Unsubscribe)
	return app
}
