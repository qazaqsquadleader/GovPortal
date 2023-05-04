package middlewares

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из заголовка запроса
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			// Если токен не был получен, возвращаем ошибку
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Проверяем валидность токена
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// В этой функции мы должны вернуть секретный ключ, используемый для создания токена
			// Обычно этот ключ хранится в настройках приложения
			return []byte("my_secret_key"), nil
		})
		if err != nil {
			// Если произошла ошибка при проверке токена, возвращаем ошибку
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Если токен валидный, сохраняем ID пользователя в контексте Gin
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["id"].(string)
			c.Set("userID", userID)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Продолжаем выполнение запроса
		c.Next()
	}
}
