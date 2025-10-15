package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Delay(delay time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		c.Next()
	}
}
