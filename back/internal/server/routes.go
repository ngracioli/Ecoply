package server

import (
	"ecoply/internal/domain/handlers"
	"ecoply/internal/domain/middlewares"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const htmlPath = "web/static"

func registerRoutes(router *gin.Engine) {
	router.LoadHTMLGlob(htmlPath + "/index.html")

	router.Use(middlewares.Cors("*"))

	router.GET("/", rootHandler)
	router.GET("/health", healthHandler)
	router.NoRoute(notFoundHandler)

	api := router.Group("api")

	v1 := api.Group("v1", middlewares.ContentType(gin.MIMEJSON))
	{
		auth := v1.Group("auth")
		{
			auth.POST("login", handlers.LoginHandler)
			auth.POST("signup", handlers.SignUpHandler)
			auth.GET("me", middlewares.JwtAuthMiddleware(), handlers.MeHandler)
			auth.POST("refresh-token", middlewares.JwtAuthMiddleware(), handlers.RefreshTokenHandler)

			available := auth.Group("available")
			{
				available.GET("", handlers.AvailabilityHandler)
			}
		}

		address := v1.Group("cnpj")
		{
			address.GET(":cnpj", handlers.CompanyByCnpj)
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
