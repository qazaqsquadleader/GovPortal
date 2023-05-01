package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckToken(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, c)
		return
	}
	// logic
}
