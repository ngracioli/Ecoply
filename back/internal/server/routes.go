package server

import (
	"ecoply/internal/domain/middlewares"
	authHandlers "ecoply/internal/domain/routes/auth"
	cnpjHandlers "ecoply/internal/domain/routes/cnpj"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("view/index.html")

	router.Use(middlewares.Cors("*"))

	router.GET("/", rootHandler)
	router.GET("/health", healthHandler)
	router.NoRoute(notFoundHandler)

	api := router.Group("api")

	v1 := api.Group("v1", middlewares.ContentType(gin.MIMEJSON))
	{
		auth := v1.Group("auth")
		{
			auth.POST("login", authHandlers.LoginHandler)
			auth.POST("signup", authHandlers.SignUpHandler)
			auth.GET("me", middlewares.JwtAuthMiddleware(), authHandlers.MeHandler)
			auth.POST("refresh-token", middlewares.JwtAuthMiddleware(), authHandlers.RefreshTokenHandler)

			available := auth.Group("available")
			{
				available.GET("", authHandlers.AvailabilityHandler)
			}
		}

		address := v1.Group("cnpj")
		{
			address.GET(":cnpj", cnpjHandlers.CompanyByCnpj)
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
