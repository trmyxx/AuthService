package service

import "Auth/internal/model"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (service *Service) AddUser(model.User) (model.User, error) {
	return model.User{}, nil
}
