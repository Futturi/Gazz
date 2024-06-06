package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/Futturi/Gaz/internal/entities"
	"github.com/Futturi/Gaz/internal/repo"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var (
	salt1 = os.Getenv("SHA_SALT")
	salt2 = os.Getenv("JWT_SALT")
)

type AuthService struct {
	repo repo.Auth
}

func NewAuthService(repo repo.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) SignUp(user entities.User) (int, error) {
	birth, err := time.Parse("2006-01-02", user.Birthday)
	if err != nil {
		return 0, err
	}
	dbUser := entities.UserForDb{
		Username: user.Username,
		Email:    user.Email,
		Password: hashPass(user.Password),
		Birthday: birth,
	}
	return a.repo.SignUp(dbUser)
}

func hashPass(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt1)))
}

func (a *AuthService) SignIn(user entities.User) (string, error) {
	user.Password = hashPass(user.Password)
	id, err := a.repo.SignIn(user)
	if err != nil {
		return "", err
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	t, err := token.SignedString([]byte(salt2))
	if err != nil {
		return "", err
	}
	return t, nil
}
