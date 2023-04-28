package handlers

import (
	"govportal/models"

	"github.com/gin-gonic/gin"
)

func MainPage(c *gin.Context) {
	firstuser := models.User{
		UserId:   1,
		Username: "Alibek",
		Status:   "Admin",
	}
	c.JSON(200, firstuser)
}
