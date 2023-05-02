package repository

import (
	"database/sql"
	"govportal/models"
	"time"
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
	SaveToken(models.User) error
	DeleteToken(string) error
	DeleteTokenById(int) error
	CheckUserByToken(string) (models.User, error)
	CheckUserByName(models.User) error
	GetUserInfo(int) (models.User, error)
}

func (a *AuthSQL) CreateUser(User models.User) error {
	stmt, err := a.db.Prepare("INSERT INTO User(username, password,email) values(?,?,?)")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(User.Username, User.Password, User.Email); err != nil {
		return err
	}
	return nil
}

func (a *AuthSQL) CheckUser(User models.User) (models.User, error) {
	var fullUser models.User
	query := `SELECT * FROM user WHERE username=$1 and password=$2`
	row := a.db.QueryRow(query, User.Username, User.Password)
	if err := row.Scan(&fullUser.UserId, &fullUser.Username, &fullUser.Password, &fullUser.Email); err != nil {
		return fullUser, err
	}
	return fullUser, nil
}

func (a *AuthSQL) CheckUserByToken(token string) (models.User, error) {
	var fullUser models.User
	var id int
	var expiresAt string
	query := `SELECT userId,expiresAt FROM user_sessions WHERE token=$1`
	err := a.db.QueryRow(query, token).Scan(&id, &expiresAt)
	if err != nil {
		return fullUser, err
	}
	query1 := `SELECT userId, username, email FROM user WHERE userId=?`
	row := a.db.QueryRow(query1, id)
	if err := row.Scan(&fullUser.UserId, &fullUser.Username, &fullUser.Email); err != nil {
		return fullUser, err
	}
	fullUser.TokenDuration, _ = time.Parse("01-02-2006 15:04:05", expiresAt)
	return fullUser, nil
}

func (a *AuthSQL) SaveToken(User models.User) error {
	stmt, err := a.db.Prepare(`INSERT INTO user_sessions(token, expiresAt,userId) values(?,?,?)`)
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(User.Token, User.TokenDuration, User.UserId); err != nil {
		return err
	}
	return nil
}

func (a *AuthSQL) DeleteToken(token string) error {
	query := `DELETE FROM userSessions WHERE token=$1`
	_, err := a.db.Exec(query, token)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthSQL) DeleteTokenById(id int) error {
	query := `DELETE FROM userSessions WHERE userId=$1`
	_, err := a.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthSQL) GetUserByEmail(Email string) (models.User, error) {
	query := "SELECT * FROM user WHERE Email = $1"
	row := a.db.QueryRow(query, Email)

	var fullUser models.User
	err := row.Scan(&fullUser.UserId, &fullUser.Username, &fullUser.Password, &fullUser.Email)
	if err != nil {
		return models.User{}, err
	}

	return fullUser, nil
}

func (a *AuthSQL) GetUserByUsername(username string) (models.User, error) {
	query := "SELECT * FROM user WHERE username = $1"
	row := a.db.QueryRow(query, username)

	var fullUser models.User
	err := row.Scan(&fullUser.UserId, &fullUser.Username, &fullUser.Password, &fullUser.Email)
	if err != nil {
		return models.User{}, err
	}

	return fullUser, nil
}

func (a *AuthSQL) GetUserInfo(UserId int) (models.User, error) {
	query := `SELECT * FROM user WHERE username=$1`
	row := a.db.QueryRow(query, UserId)
	var fullUser models.User
	err := row.Scan(&fullUser.UserId, &fullUser.Username, &fullUser.Password, &fullUser.Email)
	if err != nil {
		return models.User{}, err
	}
	return fullUser, nil
}
