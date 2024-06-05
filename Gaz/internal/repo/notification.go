package repo

import (
	"fmt"
	"github.com/Futturi/Gaz/internal/entities"
	"github.com/jmoiron/sqlx"
)

type NotificationRepo struct {
	db *sqlx.DB
}

func NewNotificationRepo(db *sqlx.DB) *NotificationRepo {
	return &NotificationRepo{db: db}
}

func (r *NotificationRepo) SendMessage(b string) ([]entities.BirthDay, error) {
	var birth []entities.BirthDay

	query := fmt.Sprintf(`SELECT u.email, (SELECT username FROM %s WHERE id = s.main_id) 
FROM %s u 
INNER JOIN %s s ON u.id = s.user_id 
WHERE s.main_id IN (SELECT id FROM %s WHERE birthdate = $1);
`, userTable, userTable, subTable, userTable)
	if err := r.db.Select(&birth, query, b); err != nil {
		return []entities.BirthDay{}, err
	}

	return birth, nil
}
