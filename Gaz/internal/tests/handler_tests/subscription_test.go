package handler_tests

import (
	"bytes"
	"errors"
	"github.com/Futturi/Gaz/internal/handler"
	"github.com/Futturi/Gaz/internal/service"
	mock_service "github.com/Futturi/Gaz/internal/service/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

func TestHandler_Subscribe(t *testing.T) {
	type mockBehavior func(s *mock_service.MockSubscription, userID float64, username string)

	tests := []struct {
		name         string
		inputBody    string
		userID       float64
		mockBehavior mockBehavior
		expectedCode int
		expectedBody string
	}{
		{
			name:      "ok",
			inputBody: `{"username": "testuser"}`,
			userID:    1,
			mockBehavior: func(s *mock_service.MockSubscription, userID float64, username string) {
				s.EXPECT().Subscribe(userID, username).Return(nil)
			},
			expectedCode: fiber.StatusOK,
			expectedBody: `{"message":"subscribed"}`,
		},
		{
			name:         "unauthorized",
			inputBody:    `{"username": "testuser"}`,
			userID:       0,
			mockBehavior: func(s *mock_service.MockSubscription, userID float64, username string) {},
			expectedCode: fiber.StatusUnauthorized,
			expectedBody: `{"message":"Unauthorized"}`,
		},
		{
			name:         "bad request",
			inputBody:    `{"invalid_json"}`,
			userID:       1,
			mockBehavior: func(s *mock_service.MockSubscription, userID float64, username string) {},
			expectedCode: fiber.StatusBadRequest,
			expectedBody: `{"message":"incorrect data"}`,
		},
		{
			name:      "internal server error",
			inputBody: `{"username": "testuser"}`,
			userID:    1,
			mockBehavior: func(s *mock_service.MockSubscription, userID float64, username string) {
				s.EXPECT().Subscribe(userID, username).Return(errors.New("some error"))
			},
			expectedCode: fiber.StatusInternalServerError,
			expectedBody: `{"message":"smth wrong in server"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Инициализация mock
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Создание mock сервиса
			mockService := mock_service.NewMockSubscription(ctrl)
			tt.mockBehavior(mockService, tt.userID, "testuser")
			servic := &service.Serivce{Subscription: mockService}
			// Create handler with the mocked service
			handler := handler.NewHandler(servic, nil)

			// Создание роутера
			app := fiber.New()

			// Определяем маршруты
			app.Post("/subscribe", func(c *fiber.Ctx) error {
				c.Locals("user", jwt.MapClaims{"id": tt.userID})
				return handler.Subscribe(c)
			})

			// Создание запроса
			req := httptest.NewRequest("POST", "/subscribe", bytes.NewBufferString(tt.inputBody))
			req.Header.Set("Content-Type", "application/json")

			// Создание ответа
			resp, _ := app.Test(req, -1)

			// Проверка кода ответа
			assert.Equal(t, tt.expectedCode, resp.StatusCode)

			// Проверка тела ответа
			body, _ := io.ReadAll(resp.Body)
			assert.JSONEq(t, tt.expectedBody, string(body))
		})
	}
}

func TestHandler_Unsubscribe(t *testing.T) {
	type mockBehavior func(s *mock_service.MockSubscription, userID float64, username string)

	tests := []struct {
		name         string
		inputBody    string
		userID       float64
		mockBehavior mockBehavior
		expectedCode int
		expectedBody string
	}{
		{
			name:      "ok",
			inputBody: `{"username": "testuser"}`,
			userID:    1,
			mockBehavior: func(s *mock_service.MockSubscription, userID float64, username string) {
				s.EXPECT().Unsubscribe(userID, username).Return(nil)
			},
			expectedCode: fiber.StatusOK,
			expectedBody: `{"message":"unsubscribed"}`,
		},
		{
			name:         "unauthorized",
			inputBody:    `{"username": "testuser"}`,
			userID:       0,
			mockBehavior: func(s *mock_service.MockSubscription, userID float64, username string) {},
			expectedCode: fiber.StatusUnauthorized,
			expectedBody: `{"message":"Unauthorized"}`,
		},
		{
			name:         "bad request",
			inputBody:    `{"invalid_json"}`,
			userID:       1,
			mockBehavior: func(s *mock_service.MockSubscription, userID float64, username string) {},
			expectedCode: fiber.StatusBadRequest,
			expectedBody: `{"message":"incorrect data"}`,
		},
		{
			name:      "internal server error",
			inputBody: `{"username": "testuser"}`,
			userID:    1,
			mockBehavior: func(s *mock_service.MockSubscription, userID float64, username string) {
				s.EXPECT().Unsubscribe(userID, username).Return(errors.New("some error"))
			},
			expectedCode: fiber.StatusInternalServerError,
			expectedBody: `{"message":"smth wrong in server"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_service.NewMockSubscription(ctrl)
			tt.mockBehavior(mockService, tt.userID, "testuser")
			servic := &service.Serivce{Subscription: mockService}
			handler := handler.NewHandler(servic, nil)

			app := fiber.New()

			app.Post("/unsubscribe", func(c *fiber.Ctx) error {
				c.Locals("user", jwt.MapClaims{"id": tt.userID})
				return handler.Unsubscribe(c)
			})
			req := httptest.NewRequest("POST", "/unsubscribe", bytes.NewBufferString(tt.inputBody))
			req.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(req, -1)

			assert.Equal(t, tt.expectedCode, resp.StatusCode)

			body, _ := io.ReadAll(resp.Body)
			assert.JSONEq(t, tt.expectedBody, string(body))
		})
	}
}
