package handlers

import (
	"govportal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainPage(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(405, c)
		return
	}

	firstuser := models.User{
		UserId:   1,
		Username: "Alibek",
		Status:   "Admin",
	}
	c.JSON(200, firstuser)
	return
}
