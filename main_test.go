package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"os"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\
// Helper Functions

func TestIndex(t *testing.T) {
	r := Router()
	Convey("Given that I go to the index page", t, func() {
		req, _ := http.NewRequest("GET", "/", nil)
		Convey("it should return a 200 status code", func() {
			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)
			So(resp.Code, ShouldEqual, http.StatusOK)
		})
	})

}