// middleware/jwt_middleware_test.go
package handler_tests

import (
	"github.com/Futturi/Gaz/internal/handler"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	salt2 = "legjliqjwoejrqfgeniowo4i3wipreq;ksdfjbgkhoiterwpq[lasdkmcvnbfjghutriowepq[als;dk,m,cvnkfjg"
)

func TestJwtMiddleware(t *testing.T) {
	type test struct {
		name         string
		token        string
		expectedCode int
		expectedBody string
	}

	tests := []test{
		{
			name:         "missing token",
			token:        "",
			expectedCode: fiber.StatusUnauthorized,
			expectedBody: `{"message":"Missing or malformed JWT"}`,
		},
		{
			name:         "invalid token",
			token:        "Bearer invalidtoken",
			expectedCode: fiber.StatusUnauthorized,
			expectedBody: `{"message":"Invalid or expired JWT"}`,
		},
		{
			name:         "valid token",
			token:        createTestJWT(t, salt2),
			expectedCode: fiber.StatusOK,
			expectedBody: ``,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()

			app.Get("/protected", handler.JwtMiddleware, func(c *fiber.Ctx) error {
				return c.SendStatus(fiber.StatusOK)
			})

			req := httptest.NewRequest("GET", "/protected", nil)
			req.Header.Set("Authorization", tt.token)

			resp, _ := app.Test(req, -1)

			assert.Equal(t, tt.expectedCode, resp.StatusCode)

		})
	}
}

// Helper function to create a test JWT token
func createTestJWT(t *testing.T, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  1,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	require.NoError(t, err)

	return "Bearer " + tokenString
}
