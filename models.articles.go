package main

type Article struct {
	ID	int `json:"id" validate:"min=1"`
	Title string `json:"title" validate:"nonzero"`
	Content string `json"content" validate:"nonzero"`
}

var articleList = []Article{
	Article{ID: 1, Title: "Article 1", Content: "Article 1 Body"},
	Article{ID: 2, Title: "Article 2", Content: "Article 2 Body"},
}

func getAllArticles() []Article {
	return articleList
}