package repo

import (
	"github.com/Futturi/Gaz/internal/entities"
	"github.com/jmoiron/sqlx"
)

const (
	userTable = "users"
	subTable  = "subscribers"
)

type Repository struct {
	Auth
	Subscription
	Users
	Notification
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Auth: NewAuthRepo(db), Subscription: NewSubscriptionRepo(db), Users: NewUsersRepo(db), Notification: NewNotificationRepo(db)}
}

type Auth interface {
	SignUp(user entities.UserForDb) (int, error)
	SignIn(user entities.User) (int, error)
}

type Subscription interface {
	Subscribe(id float64, username string) error
	Unsubscribe(id float64, username string) error
}

type Users interface {
	GetUsers(id float64) ([]entities.UserForDb, error)
}

type Notification interface {
	SendMessage(b string) ([]entities.BirthDay, error)
}
