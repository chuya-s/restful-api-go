package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/chuya-s/restful-api-go/internal/article"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Env struct {
	db *sql.DB
}

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoit Hit")
}

func (env *Env) allArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	articles, err := article.GetAll(env.db)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(articles)
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST article Endpoit Hit")
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func handleRequests(env *Env) {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)

	myRouter.HandleFunc("/articles", env.allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", postArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("main start.")

	db, err := sql.Open("mysql", "test_user:password@tcp(0.0.0.0:3306)/test_database")
	if err != nil {
		fmt.Println("Db open error:")
	}

	if err != nil {
		fmt.Println("DB open error.")
	}

	env := &Env{db: db}

	handleRequests(env)
	fmt.Println("main end.")
}
