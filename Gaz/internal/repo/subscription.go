package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type SubscriptionRepo struct {
	db *sqlx.DB
}

func NewSubscriptionRepo(db *sqlx.DB) *SubscriptionRepo {
	return &SubscriptionRepo{db: db}
}

func (r *SubscriptionRepo) Subscribe(id float64, username string) error {
	query := fmt.Sprintf("INSERT INTO %s(user_id, main_id) SELECT $1, id FROM %s WHERE username = $2", subTable, userTable)
	_, err := r.db.Exec(query, id, username)
	slog.Info("subscribe", "id", id, "username", username)
	return err
}

func (r *SubscriptionRepo) Unsubscribe(id float64, username string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND main_id = (SELECT id FROM %s WHERE username = $2)", subTable, userTable)

	_, err := r.db.Exec(query, id, username)
	slog.Info("unsubscribe", "id", id, "username", username)
	return err
}
