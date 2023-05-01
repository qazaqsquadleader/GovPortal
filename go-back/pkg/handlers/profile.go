package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, c)
		return
	}
	CheckUserByToken(c.Request.Cookie("token"))
}
