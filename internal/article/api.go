package article

import (
	"github.com/chuya-s/restful-api-go/mysql"
)

type Article struct {
	Id      int
	Title   string
	Content string
}

func GetAll() []Article {
	queried, err := mysql.Db.Query("select * from article")
	if err != nil {
		panic(err.Error())
	}
	articles := []Article{}
	for queried.Next() {
		article := Article{}
		err = queried.Scan(&article.Id, &article.Title, &article.Content)
		if err != nil {
			panic(err.Error())
		}
		articles = append(articles, article)
		queried.Close()
	}
	return articles
}
