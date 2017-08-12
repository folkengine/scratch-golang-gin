package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.template",
			gin.H{
				"title": "Home Page",
			},
		)
	})
	router.Run()
}