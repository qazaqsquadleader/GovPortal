package service

import (
	"govportal/models"
	"govportal/pkg/repository"
)

type AuthServiceStruct struct {
	repo repository.IAuthSQL
}

type AuthService interface {
	CreateUser(models.User) error
	GetUserByToken(string) (models.User, error)
	CheckUser(models.User) (models.User, error)
	DeleteToken(string) error
	PersonalInfo(int) (models.User, error)
}

func NewAuthService(repoAuth repository.IAuthSQL) AuthService {
	return &AuthServiceStruct{
		repoAuth,
	}
}

func (a *AuthServiceStruct) GetUserByToken(token string) (models.User, error) {
	var user models.User
	user, err := a.repo.CheckUserByToken(token)
	if err != nil {
		return user, err
	}
	return user, nil
}

// func (a *AuthServiceStruct) CreateUser(user models.User) error {
// 	// if err := isValidEmail(user.Email); err != nil {
// 	// 	return project_error.NewUserError(err.Error(), http.StatusBadRequest)
// 	// }
// 	// if err := isValidUsername(user.Username); err != nil {
// 	// 	return project_error.NewUserError(err.Error(), http.StatusBadRequest)
// 	// }
// 	// if err := isValidPassword(user.Password); err != nil {
// 	// 	return project_error.NewUserError(err.Error(), http.StatusBadRequest)
// 	// }
// 	// _, err := a.repo.GetUserByUsername(user.Username)
// 	// if err == nil {
// 	// 	return project_error.NewUserError("username already exists", http.StatusConflict)
// 	// } else if !errors.Is(err, sql.ErrNoRows) {
// 	// 	return project_error.NewServerError(err.Error())
// 	// }

// 	// _, err = a.repo.GetUserByEmail(user.Email)
// 	// if err == nil {
// 	// 	return project_error.NewUserError("email already exists", http.StatusConflict)
// 	// } else if !errors.Is(err, sql.ErrNoRows) {
// 	// 	return project_error.NewServerError(err.Error())
// 	// }

// 	// if err := a.repo.CreateUser(user); err != nil {
// 	// 	return project_error.NewServerError(err.Error())
// 	// }
// 	return nil
// }
