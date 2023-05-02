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
	CheckUser(models.User) (models.User, error)
}

func (a *AuthServiceStruct) GetUserByToken(token string) (models.User, error) {
	var user models.User
	user, err := a.repo.CheckUserByToken(token)
	if err != nil {
		return user, err
	}
	return user, nil
}

func CheckUser(user models.User) (models.User, error) {
	// if err := isValidUsername(user.Username); err != nil {
	// 	return user, errormsg.NewUserError("invalid username", http.StatusBadRequest)
	// }
	// if err := isValidPassword(user.Password); err != nil {
	// 	return user, errormsg.NewUserError("invalid password", http.StatusBadRequest)
	// }
	// user, err := a.repository.CheckUser(user)
	// if err != nil {
	// 	return user, errormsg.NewUserError("User does not exist", http.StatusUnauthorized)
	// }
	// token, err := uuid.NewV4()
	// if err != nil {
	// 	return user, errormsg.NewInternalError("error generating token")
	// }
	// user.Token = token.String()
	// user.TokenDuration = time.Now().Add(72 * time.Hour)
	// a.repo.DeleteTokenById(user.ID)
	// if err := a.repo.SaveToken(user); err != nil {
	// 	return user, errormsg.NewInternalError("error saving token to database")
	// }
	return user, nil
}
