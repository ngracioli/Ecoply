package middlewares

import "github.com/gin-gonic/gin"

func ContentType(contentType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", contentType)
		c.Next()
	}
}
