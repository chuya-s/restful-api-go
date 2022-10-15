package article

import (
	"database/sql"
)

type Article struct {
	Id      int
	Title   string
	Content string
}

func GetAll(db *sql.DB) ([]Article, error) {
	queried, err := db.Query("select * from article;")
	if err != nil {
		return nil, err
	}
	defer queried.Close()
	articles := []Article{}
	for queried.Next() {
		article := Article{}
		err = queried.Scan(&article.Id, &article.Title, &article.Content)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}
