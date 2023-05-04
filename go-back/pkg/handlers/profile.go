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

	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "Failed to get user info"})
	// 	return
	// }

	c.JSON(http.StatusOK, c)
}
