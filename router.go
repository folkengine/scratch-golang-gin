package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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