package server

import (
	domainAuth "ecoply/internal/domain/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("view/index.html")

	router.GET("/", rootHandler)
	router.GET("/health", healthHandler)
	router.NoRoute(notFoundHandler)

	api := router.Group("api")

	v1 := api.Group("v1", Cors("*"), ContentType("application/json; charset=utf-8"))
	{
		auth := v1.Group("auth")
		{
			auth.POST("login", domainAuth.LoginHandler)
			auth.POST("signup", domainAuth.SignUpHandler)
		}
	}
}

func rootHandler(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Ecoply API",
	})
}

func notFoundHandler(c *gin.Context) {
	var accept string = c.Request.Header.Get("Accept")

	if strings.Contains(accept, "application/json") {
		c.JSON(404, gin.H{
			"message": "Not Found",
		})
		return
	}

	c.Redirect(302, "/")
}

func healthHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
