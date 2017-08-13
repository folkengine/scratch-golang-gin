package main

import (
	_ "github.com/gin-gonic/gin"
	_ "github.com/drone/drone/router"
)

func main() {
	router := Router()
	initializeRoutes(router)
	router.Run()
}