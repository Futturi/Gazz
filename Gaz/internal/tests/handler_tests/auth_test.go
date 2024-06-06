package handler_tests

import (
	"bytes"
	"github.com/Futturi/Gaz/internal/entities"
	"github.com/Futturi/Gaz/internal/handler"
	"github.com/Futturi/Gaz/internal/service"
	mock_service "github.com/Futturi/Gaz/internal/service/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_SignUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuth, user entities.User)

	tests := []struct {
		name                string
		inputBody           string
		inputUser           entities.User
		mockBehavior        mockBehavior
		expectedCode        int
		expectedRequestBody string
	}{
		{
			name:      "ok",
			inputBody: `{"username": "username", "password": "password", "email": "email@r.com", "birthday": "2006-01-02"}`,
			inputUser: entities.User{
				Username: "username",
				Password: "password",
				Email:    "email@r.com",
				Birthday: "2006-01-02",
			},
			mockBehavior: func(s *mock_service.MockAuth, user entities.User) {
				s.EXPECT().SignUp(user).Return(1, nil)
			},
			expectedCode:        200,
			expectedRequestBody: `{"id": 1}`,
		},
		{
			name:      "bad request",
			inputBody: `{"legjlkqwk4":"wdlkjgwe"}`,
			inputUser: entities.User{},
			mockBehavior: func(s *mock_service.MockAuth, user entities.User) {
			},
			expectedCode:        400,
			expectedRequestBody: `{"message": "invalid input body"}`,
		},
		{
			name:      "error with email",
			inputBody: `{"username": "username", "password": "password", "email": "emailr.com", "birthday": "2006-01-02"}`,
			inputUser: entities.User{},
			mockBehavior: func(s *mock_service.MockAuth, user entities.User) {
			},
			expectedCode:        400,
			expectedRequestBody: `{"message": "some error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Инициализируем mock
			c := gomock.NewController(t)
			defer c.Finish()
			mock := mock_service.NewMockAuth(c)
			test.mockBehavior(mock, test.inputUser)
			// Инициализируем хэндл
			service := &service.Serivce{Auth: mock}

			handle := handler.NewHandler(service, nil)
			// Создаем роутер и маршруты
			r := fiber.New()
			r.Post("/auth/signup", handle.SignUp)
			// Создаем запрос
			req := httptest.NewRequest("POST", "/auth/signup", bytes.NewBufferString(test.inputBody))
			req.Header.Set("Content-Type", "application/json")
			// Получаем ответ
			resp, _ := r.Test(req, -1)
			// Проверяем ответ
			assert.Equal(t, test.expectedCode, resp.StatusCode)

		})
	}
}

func TestHandler_SignIn(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuth, user entities.User)
	tests := []struct {
		name                string
		inputBody           string
		inputUser           entities.User
		mockBehavior        mockBehavior
		expectedCode        int
		expectedRequestBody string
	}{
		{
			name:      "ok",
			inputBody: `{"email": "email@r.com", "password": "password"}`,
			inputUser: entities.User{
				Email:    "email@r.com",
				Password: "password",
			},
			mockBehavior: func(s *mock_service.MockAuth, user entities.User) {
				s.EXPECT().SignIn(user).Return("token", nil)
			},
			expectedCode:        200,
			expectedRequestBody: `{"token": "token"}`,
		},
		{
			name:      "bad request",
			inputBody: `{"213241": "12421", "password": "12}`,
			inputUser: entities.User{},
			mockBehavior: func(s *mock_service.MockAuth, user entities.User) {
			},
			expectedCode:        400,
			expectedRequestBody: `{"message": "invalid input body"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			mock := mock_service.NewMockAuth(c)
			test.mockBehavior(mock, test.inputUser)

			service := &service.Serivce{Auth: mock}

			handle := handler.NewHandler(service, nil)

			r := fiber.New()
			r.Post("/auth/signin", handle.SignIn)
			req := httptest.NewRequest("POST", "/auth/signin", bytes.NewBufferString(test.inputBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := r.Test(req, -1)
			assert.Equal(t, test.expectedCode, resp.StatusCode)
		})
	}
}
