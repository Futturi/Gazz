package service

import (
	"github.com/Futturi/Gaz/internal/entities"
	"github.com/Futturi/Gaz/internal/repo"
	"time"
)

type NotificationService struct {
	repo repo.Notification
}

func NewNotificationService(repo repo.Notification) *NotificationService {
	return &NotificationService{repo: repo}
}

func (a *NotificationService) SendMessage() ([]entities.BirthDay, error) {
	b := time.Now().Add(time.Hour * 24).Format("2006-01-02")
	return a.repo.SendMessage(b)
}
