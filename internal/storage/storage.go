package storage

import (
	"github.com/trmyxx/AuthService/initializers"
	"github.com/trmyxx/AuthService/internal/model"
	"gorm.io/gorm"
)

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

func (storage *Storage) AddUser(body model.User, hash string) *gorm.DB {
	user := model.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)
	return result
}
