package handlers

import (
	"govportal/models"
	"govportal/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	// 	services *service.Service
	Router *http.ServeMux
	Logger *logger.Logger
}

func SignIn(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, c)
		return
	}
	var myStruct models.User
	if err := c.BindJSON(myStruct); err != nil {
		c.JSON(http.StatusBadRequest, c)
		return
	}
	// .CheckUserByName(myStruct)
}
