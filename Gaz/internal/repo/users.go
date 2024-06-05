package repo

import (
	"fmt"
	"github.com/Futturi/Gaz/internal/entities"
	"github.com/jmoiron/sqlx"
)

type UsersRepo struct {
	db *sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) GetUsers(id float64) ([]entities.UserForDb, error) {
	var users []entities.UserForDb
	query := fmt.Sprintf("SELECT username, birthdate FROM %s WHERE id != $1", userTable)
	if err := r.db.Select(&users, query, id); err != nil {
		return nil, err
	}
	return users, nil
}
