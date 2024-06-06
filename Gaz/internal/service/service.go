package service

import (
	"github.com/Futturi/Gaz/internal/entities"
	"github.com/Futturi/Gaz/internal/repo"
)

type Serivce struct {
	Auth
	Subscription
	Users
	Notification
}

func NewSerivce(repo *repo.Repository) *Serivce {
	return &Serivce{Auth: NewAuthService(repo.Auth),
		Subscription: NewSubscriptionService(repo.Subscription),
		Users:        NewUsersService(repo.Users),
		Notification: NewNotificationService(repo.Notification)}
}

//go:generate mockgen -source=service.go -destination=mocks/mockr.go
type Auth interface {
	SignUp(user entities.User) (int, error)
	SignIn(user entities.User) (string, error)
}

type Subscription interface {
	Subscribe(id float64, username string) error
	Unsubscribe(id float64, username string) error
}

type Users interface {
	GetUsers(id float64) ([]entities.User, error)
}

type Notification interface {
	SendMessage() ([]entities.BirthDay, error)
}
