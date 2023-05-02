package handlers

import (
	"govportal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, c)
		return
	}
	userInfo, err := repo.AuthService.PersonalInfo()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to get user info"})
		return
	}
	// Преобразуем информацию о пользователе в нужный формат для вывода на страницу
	profileInfo := struct {
		Username string
		Email    string
	}{
		Username: userInfo.(models.User).Username,
		Email:    userInfo.(models.User).Email,
	}
	// Выводим информацию на страницу
	c.JSON(http.StatusOK, profileInfo)
}
