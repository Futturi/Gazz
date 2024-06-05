package service

import (
	"github.com/Futturi/Gaz/internal/entities"
	"github.com/Futturi/Gaz/internal/repo"
)

type UsersService struct {
	repo repo.Users
}

func NewUsersService(repo repo.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (a *UsersService) GetUsers(id float64) ([]entities.User, error) {
	user, err := a.repo.GetUsers(id)
	if err != nil {
		return []entities.User{}, err
	}
	newUs := make([]entities.User, 0)
	for _, v := range user {
		newUs = append(newUs, entities.User{
			Username: v.Username,
			Birthday: v.Birthday.Format("2006-01-02"),
		})
	}
	return newUs, nil
}
