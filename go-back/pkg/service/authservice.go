package service

import (
	"govportal/models"
	"govportal/pkg/repository"
)

type AuthServiceStruct struct {
	repo repository.IAuthSQL
}

type AuthService interface {
	CreateUser(models.User, error)
	CheckUserByToken(string) (models.User, error)
}

func (a *AuthServiceStruct) GetUserByToken(token string) (models.User, error) {
	var user models.User
	user, err := a.repo.CheckUserByToken(token)
	if err != nil {
		return user, err
	}
	return user, nil
}
