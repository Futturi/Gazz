// handlers/get_users_test.go
package handler_tests

import (
	"errors"
	"github.com/Futturi/Gaz/internal/entities"
	"github.com/Futturi/Gaz/internal/handler"
	"github.com/Futturi/Gaz/internal/service"
	"io"
	"net/http/httptest"
	"testing"

	mock_service "github.com/Futturi/Gaz/internal/service/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetUsers(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUsers, userID float64)

	tests := []struct {
		name         string
		userID       float64
		mockBehavior mockBehavior
		expectedCode int
		expectedBody string
	}{
		{
			name:   "ok",
			userID: 1,
			mockBehavior: func(s *mock_service.MockUsers, userID float64) {
				users := []entities.User{
					{Username: "user1", Email: "user1@example.com"},
					{Username: "user2", Email: "user2@example.com"},
				}
				s.EXPECT().GetUsers(userID).Return(users, nil)
			},
			expectedCode: fiber.StatusOK,
			expectedBody: `{"users":[{"birthday":"","username":"user1","password":"","email":"user1@example.com"},{"birthday":"","username":"user2","password":"","email":"user2@example.com"}]}`,
		},
		{
			name:         "unauthorized",
			userID:       0,
			mockBehavior: func(s *mock_service.MockUsers, userID float64) {},
			expectedCode: fiber.StatusUnauthorized,
			expectedBody: `{"message":"Unauthorized"}`,
		},
		{
			name:   "internal server error",
			userID: 1,
			mockBehavior: func(s *mock_service.MockUsers, userID float64) {
				s.EXPECT().GetUsers(userID).Return(nil, errors.New("some error"))
			},
			expectedCode: fiber.StatusInternalServerError,
			expectedBody: `{"message":"smth wrong in server"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Инициализируем mock сервис
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := mock_service.NewMockUsers(ctrl)
			tt.mockBehavior(mockService, tt.userID)

			// Инициализируем хэндлер
			servic := &service.Serivce{Users: mockService}

			handler := handler.NewHandler(servic, nil)

			// Создаем роутер
			app := fiber.New()

			app.Get("/users", func(c *fiber.Ctx) error {
				c.Locals("user", jwt.MapClaims{"id": tt.userID})
				return handler.GetUsers(c)
			})

			// Создаем запрос
			req := httptest.NewRequest("GET", "/users", nil)
			req.Header.Set("Content-Type", "application/json")

			// Отправляем запрос
			resp, _ := app.Test(req, -1)

			// Проверяем код ответа
			assert.Equal(t, tt.expectedCode, resp.StatusCode)

			// Проверяем тело ответа
			body, _ := io.ReadAll(resp.Body)
			assert.JSONEq(t, tt.expectedBody, string(body))
		})
	}
}
