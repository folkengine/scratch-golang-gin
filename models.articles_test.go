package main

import (
	"testing"
	"fmt"
)

func TestGetAllArticles(t *testing.T) {
	articles := getAllArticles()
	for count, article := range articles {
		title := fmt.Sprintf("Article %d", count)
		article.Title = title
	}
}
