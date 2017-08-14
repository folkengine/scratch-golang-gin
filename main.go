package main

import (
	_ "github.com/gin-gonic/gin"
)

func main() {
	router := Router()
	initializeRoutes(router)
	router.Run()
}
