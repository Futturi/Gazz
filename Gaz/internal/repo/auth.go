package repo

import (
	"fmt"
	"github.com/Futturi/Gaz/internal/entities"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) SignUp(user entities.UserForDb) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(username, password, email, birthdate) VALUES ($1, $2, $3, $4) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Username, user.Password, user.Email, user.Birthday)
	if err := row.Scan(&id); err != nil {
		slog.Error("error with query", "error", err)
		return 0, err
	}
	slog.Info("new user in db", "id", id)
	return id, nil
}

func (r *AuthRepo) SignIn(user entities.User) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE email = $1 AND password = $2", userTable)
	row := r.db.QueryRow(query, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
