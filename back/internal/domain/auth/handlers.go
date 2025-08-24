package auth

import "github.com/gin-gonic/gin"

func LoginHandler(c *gin.Context) {
	c.Data(204, "", nil)
}

func SignUpHandler(c *gin.Context) {
	c.Data(204, "", nil)
}
