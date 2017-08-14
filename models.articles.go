package main

type article struct {
	ID	int `json:"id"`
	Title string `json:"title"`
	Content string `json"content"`
}

var articleList = []article {
	article{ID: 1, Title: "Article 1", Content: "Article 1 Body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 Body"},
}

func getAllArticles() []article {
	return articleList
}