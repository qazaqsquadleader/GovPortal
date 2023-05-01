package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, c)
		errors.New("Error method not allowed")
		return
	}
	// Logic
}
