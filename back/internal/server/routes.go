package server

import (
	"ecoply/internal/domain/middlewares"
	domainAuth "ecoply/internal/domain/routes/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("view/index.html")

	router.GET("/", rootHandler)
	router.GET("/health", healthHandler)
	router.NoRoute(notFoundHandler)

	api := router.Group("api")

	v1 := api.Group("v1", middlewares.Cors("*"), middlewares.ContentType(gin.MIMEJSON))
	{
		auth := v1.Group("auth")
		{
			auth.POST("login", domainAuth.LoginHandler)
			auth.POST("signup", domainAuth.SignUpHandler)
			auth.GET("me", middlewares.JwtAuthMiddleware(), domainAuth.MeHandler)

			available := auth.Group("available")
			{
				available.GET("", domainAuth.AvailabilityHandler)
			}
		}
	}
}

func rootHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Ecoply API",
	})
}

func notFoundHandler(c *gin.Context) {
	var accept string = c.Request.Header.Get("Accept")

	if strings.Contains(accept, gin.MIMEJSON) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
		return
	}

	c.Redirect(http.StatusFound, "/")
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
