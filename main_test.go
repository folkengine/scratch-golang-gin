package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"os"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	router = Router()
	initializeRoutes(router)
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func TestIndex(t *testing.T) {
	Convey("Given that I go to the index page", t, func() {
		req, _ := http.NewRequest("GET", "/", nil)
		Convey("it should return a 200 status code", func() {
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			So(resp.Code, ShouldEqual, http.StatusOK)
		})
	})
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\
// Helper Functions