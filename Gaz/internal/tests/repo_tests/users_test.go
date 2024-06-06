// repository/users_repo_test.go
package repo_tests

import (
	"fmt"
	"github.com/Futturi/Gaz/internal/entities"
	"github.com/Futturi/Gaz/internal/repo"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestUsersRepo_GetUsers(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	require.NoError(t, err)
	defer db.Close()

	r := repo.NewUsersRepo(db)

	userID := float64(1)

	users := []entities.UserForDb{
		{Username: "user1", Birthday: time.Now()},
		{Username: "user2", Birthday: time.Now()},
	}

	tests := []struct {
		name          string
		mockBehavior  func()
		expectedUsers []entities.UserForDb
		expectedErr   error
	}{
		{
			name: "successful retrieval",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"username", "birthdate"}).
					AddRow(users[0].Username, users[0].Birthday).
					AddRow(users[1].Username, users[1].Birthday)
				mock.ExpectQuery("SELECT username, birthdate FROM users WHERE id !=").
					WithArgs(userID).
					WillReturnRows(rows)
			},
			expectedUsers: users,
			expectedErr:   nil,
		},
		{
			name: "error in query",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT username, birthdate FROM users WHERE id !=").
					WithArgs(userID).
					WillReturnError(fmt.Errorf("query error"))
			},
			expectedUsers: nil,
			expectedErr:   fmt.Errorf("query error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			retrievedUsers, err := r.GetUsers(userID)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedUsers, retrievedUsers)
			}

			// Ensure all expectations are met
			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}
