package repository

import (
	"database/sql"
	"errors"
	"govportal/models"
)

func NewAuthSQL(db *sql.DB) *AuthSQL {
	return &AuthSQL{
		db: db,
	}
}

type AuthSQL struct {
	db *sql.DB
}

type IAuthSQL interface {
	CreateUser(models.User) error
	CheckUser(models.User) (models.User, error)
	CheckUserByToken(string) (models.User, error)
}

func (a *AuthSQL) CreateUser(user models.User) error {
	stmt, err := a.db.Prepare("INSERT INTO User(username, password,email) values(?,?,?)")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(user.Username, user.Password, user.Email); err != nil {
		return errors.New("")
	}
	return nil
}

func (a *AuthSQL) CheckUser(user models.User) (models.User, error) {
	var fulUser models.User
	query := `SELECT * FROM user WHERE username=$1 and password=$2`
	row := a.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&fulUser.UserId, &fulUser.Username, &fulUser.Password, &fulUser.Email); err != nil {
		return fulUser, err
	}
	return fulUser, nil
}
