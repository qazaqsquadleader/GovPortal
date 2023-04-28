package handlers

import (
	"database/sql"
	"govportal/models"
	"govportal/pkg/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthSQL struct {
	db *sql.DB
}

type IAuthSQL interface {
	CreateUser(models.User) error
}

func NewAuthSQL(db *sql.DB) *AuthSQL {
	return &AuthSQL{
		db: db,
	}
}

func SignUp(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal Server Error"})
		return
	}
	var a repository.AuthSQL
	a.CreateUser(newUser)
	c.String(200, "Hello World!")
	c.JSON(http.StatusCreated, gin.H{"message": "User registered succesfully"})
}
