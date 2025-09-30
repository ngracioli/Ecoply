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

		offer := v1.Group("offers").Use(
			middlewares.JwtAuthMiddleware(),
		)
		{
			offer.GET(":uuid", handlers.OfferByUuidHanlder)
			offer.GET("", handlers.OfferListHandler)
			offer.POST("", middlewares.SupplierMiddleware(), handlers.CreateOfferHandler)
			offer.PUT(":uuid", middlewares.SupplierMiddleware(), handlers.UpdateOfferHandler)
			offer.DELETE(":uuid", middlewares.SupplierMiddleware(), handlers.DeleteOfferHandler)
		}

		me := v1.Group("me").Use(middlewares.JwtAuthMiddleware())
		{
			me.GET("", handlers.MeHandler)
			me.GET("offers", middlewares.SupplierMiddleware(), handlers.UserOffersHandler)
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
