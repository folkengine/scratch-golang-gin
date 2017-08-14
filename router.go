package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Router returns a Gin Engine Router loaded with the templates.
func Router() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	return router
}

func initializeRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.template",
			gin.H{
				"title": "Home Page",
			},
		)
	})
}
