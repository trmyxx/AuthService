package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/trmyxx/AuthService/internal/model"
	"github.com/trmyxx/AuthService/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) *Service {
	return &Service{storage: storage}
}

func (service *Service) Signup(body model.User) error {
	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		return err
	}

	result := service.storage.AddUser(body, string(hash))

	if result.Error != nil {
		return result.Error
	}

	//Respond
	return nil
}

func (service *Service) Login(body model.User) (string, error) {
	//Look up requested user
	user := service.storage.GetUser(body)
	if user.ID == 0 {
		return "", errors.New("Invalid email")
	}

	//Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return "", errors.New("Invalid password")
	}

	//Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", errors.New("Failed to create jwt")
	}

	//Send it back
	return tokenString, nil
}
