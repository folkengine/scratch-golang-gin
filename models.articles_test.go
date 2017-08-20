package main

import (
	"fmt"
	"testing"
	"gopkg.in/validator.v2"
)

func TestGetAllArticles(t *testing.T) {
	articles := getAllArticles()
	for count, article := range articles {
		title := fmt.Sprintf("Article %d", count)
		article.Title = title

		if errs := validator.Validate(article); errs != nil {
			t.Fail()
		}
	}
}

func TestValidate_InvalidArticle(t *testing.T) {
	nur := Article{ID: 0, Title: "Article 0", Content: "Article 1 Body"}
	if errs := validator.Validate(nur); errs == nil {
		t.Fail()
	}
}

