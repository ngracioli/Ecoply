package server

import "github.com/gin-gonic/gin"

func registerRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("view/index.html")

	router.GET("/", rootHandler)
	router.GET("/health", healthHandler)
	router.NoRoute(notFoundHandler)

	// api := router.Group("api")

	// v1 := api.Group("v1", Cors("*"), ContentType("application/json; charset=utf-8"))
	// {
	// }
}

func rootHandler(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Ecoply API",
	})
}

func notFoundHandler(c *gin.Context) {
	c.Redirect(302, "/")
}

func healthHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
