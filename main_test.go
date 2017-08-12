package main

import (
	"testing"
	"github.com/gin-gonic/gin"
	"os"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

func TestIndex(t *testing.T) {
	r := getRouter(true)
	if r == nil {
		t.Fail()
	}
}