package server

import (
	"ecoply/internal/domain/handlers"
	"ecoply/internal/domain/middlewares"
	"ecoply/internal/domain/services"
	"time"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const htmlPath = "web/static"

func registerRoutes(router *gin.Engine, s *ServerContext) {
	var jwtService = services.NewJwtService(s.Cfg)
	var authHandlers handlers.AuthHandlers = s.Handlers.AuthHandlers
	var offerHandlers handlers.OfferHandlers = s.Handlers.OfferHandlers
	var cnpjHandlers handlers.CnpjHandlers = s.Handlers.CnpjHandlers
	var cceeHandlers handlers.CceeHandlers = s.Handlers.CceeHandlers
	var purchaseHandlers handlers.PurchaseHandlers = s.Handlers.PurchaseHandlers
	var contractHandlers handlers.ContractHandlers = s.Handlers.ContractHandlers
	var analyticsHandlers handlers.AnalyticsHandlers = s.Handlers.AnalyticsHandlers

	router.LoadHTMLGlob(htmlPath + "/index.html")

	router.Use(
		middlewares.Cors("*"),
		middlewares.Delay(time.Duration(s.Cfg.ServerDelay)),
	)

	router.GET("/", rootHandler)
	router.GET("/health", healthHandler)
	router.NoRoute(notFoundHandler)

	api := router.Group("api")

	v1 := api.Group("v1", middlewares.ContentType(gin.MIMEJSON))
	{
		auth := v1.Group("auth")
		{
			auth.POST("login", authHandlers.Login)
			auth.POST("signup", authHandlers.SignUp)
			auth.POST("refresh-token", middlewares.JwtAuthMiddleware(
				s.Services.UserService,
				jwtService,
			), authHandlers.RefreshToken)

			available := auth.Group("available")
			{
				available.GET("", authHandlers.Availability)
			}
		}

		address := v1.Group("cnpj")
		{
			address.GET(":cnpj", cnpjHandlers.CompanyByCnpj)
		}

		ccee := v1.Group("ccee")
		{
			ccee.GET("agents/:cnpj", cceeHandlers.GetAgentsByCnpj)
		}

		offer := v1.Group("offers", middlewares.JwtAuthMiddleware(
			s.Services.UserService,
			jwtService,
		))
		{
			offer.GET(":uuid", offerHandlers.FindByUuid)
			offer.GET("", offerHandlers.List)
			offer.POST("", middlewares.SupplierMiddleware(s.Services.UserTypeService), offerHandlers.Create)
			offer.PUT(":uuid", middlewares.SupplierMiddleware(s.Services.UserTypeService), offerHandlers.Update)
			offer.DELETE(":uuid", middlewares.SupplierMiddleware(s.Services.UserTypeService), offerHandlers.Delete)

			purchase := offer.Group(":uuid/purchases")
			{
				purchase.GET("", middlewares.SupplierMiddleware(s.Services.UserTypeService), offerHandlers.Purchases)
				purchase.POST("", purchaseHandlers.Create)
			}

			checkout := offer.Group(":uuid/checkout")
			{
				checkout.GET("")
			}
		}

		purchases := v1.Group("purchases", middlewares.JwtAuthMiddleware(
			s.Services.UserService,
			jwtService,
		))
		{
			purchases.GET("", purchaseHandlers.ListPurchases)
			purchases.GET(":uuid", purchaseHandlers.FindByUuid)
			purchases.POST(":uuid/cancel", purchaseHandlers.Cancel)
			purchases.GET(":uuid/contract", contractHandlers.Get)
		}

		sales := v1.Group("sales", middlewares.JwtAuthMiddleware(
			s.Services.UserService,
			jwtService,
		), middlewares.SupplierMiddleware(s.Services.UserTypeService))
		{
			sales.GET("", purchaseHandlers.ListSales)
		}

		me := v1.Group("me").Use(middlewares.JwtAuthMiddleware(
			s.Services.UserService,
			jwtService,
		))
		{
			me.GET("", authHandlers.Me)
			me.GET("offers", middlewares.SupplierMiddleware(s.Services.UserTypeService), offerHandlers.FromUser)
			me.GET("analytics", analyticsHandlers.User)
		}

		analytics := v1.Group("analytics")
		{
			analytics.GET("platform", analyticsHandlers.Platform)
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
