package service

import "github.com/Futturi/Gaz/internal/repo"

type SubscriptionService struct {
	repo repo.Subscription
}

func NewSubscriptionService(repo repo.Subscription) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (a *SubscriptionService) Subscribe(id float64, username string) error {
	return a.repo.Subscribe(id, username)
}

func (a *SubscriptionService) Unsubscribe(id float64, username string) error {
	return a.repo.Unsubscribe(id, username)
}
