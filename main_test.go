package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/PuerkitoBio/goquery"
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
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		Convey("it should return a 200 status code", func() {
			So(resp.Code, ShouldEqual, http.StatusOK)
		})
		Convey("it should have the title of Home Page", func() {
			So(getField(resp.Body, "title"), ShouldEqual, "Home Page")
		})
	})
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\
// Helper Functions

func getField(body *bytes.Buffer, fieldname string) string {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(body.String()))
	return doc.Find(fieldname).Text()
}
