package middlewares

import "github.com/gin-gonic/gin"

func BasicAuth(f func(c *gin.Context)) bool {
	return false
}
