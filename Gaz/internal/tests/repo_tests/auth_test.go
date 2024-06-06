// repository/auth_repo_test.go
package repo_tests

import (
	"fmt"
	"github.com/Futturi/Gaz/internal/entities"
	"github.com/Futturi/Gaz/internal/repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"time"

	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func TestAuthRepo_SignUp(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	require.NoError(t, err)
	defer db.Close()
	r := repo.NewAuthRepo(db)
	user := entities.UserForDb{
		Username: "username",
		Password: "password",
		Email:    "email@domain.com",
		Birthday: time.Now(),
	}

	tests := []struct {
		name         string
		mockBehavior func()
		expectedID   int
		expectedErr  error
	}{
		{
			name: "successful sign up",
			mockBehavior: func() {
				mock.ExpectQuery("INSERT INTO users").
					WithArgs(user.Username, user.Password, user.Email, user.Birthday).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			expectedID:  1,
			expectedErr: nil,
		},
		{
			name: "error in query",
			mockBehavior: func() {
				mock.ExpectQuery("INSERT INTO users").
					WithArgs(user.Username, user.Password, user.Email, user.Birthday).
					WillReturnError(fmt.Errorf("query error"))
			},
			expectedID:  0,
			expectedErr: fmt.Errorf("query error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			id, err := r.SignUp(user)

			assert.Equal(t, tt.expectedID, id)
			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			// Ensure all expectations are met
			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}

func TestAuthRepo_SignIn(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	require.NoError(t, err)
	defer db.Close()

	r := repo.NewAuthRepo(db)

	user := entities.User{
		Email:    "email@domain.com",
		Password: "password",
	}

	tests := []struct {
		name         string
		mockBehavior func()
		expectedID   int
		expectedErr  error
	}{
		{
			name: "successful sign in",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT id FROM users WHERE email = \\$1 AND password = \\$2").
					WithArgs(user.Email, user.Password).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			expectedID:  1,
			expectedErr: nil,
		},
		{
			name: "user not found",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT id FROM users WHERE email = \\$1 AND password = \\$2").
					WithArgs(user.Email, user.Password).
					WillReturnError(fmt.Errorf("sql: no rows in result set"))
			},
			expectedID:  0,
			expectedErr: fmt.Errorf("sql: no rows in result set"),
		},
		{
			name: "error in query",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT id FROM users WHERE email = \\$1 AND password = \\$2").
					WithArgs(user.Email, user.Password).
					WillReturnError(fmt.Errorf("query error"))
			},
			expectedID:  0,
			expectedErr: fmt.Errorf("query error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			id, err := r.SignIn(user)

			assert.Equal(t, tt.expectedID, id)
			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			// Ensure all expectations are met
			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}
