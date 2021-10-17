package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func main() {
	// database data
	Articles = []Article{
		{Id: "1", Title: "Model Predictive Control", Desc: "MPC Book", Content: "Model predictive control"},
		{Id: "2", Title: "Convex Optimization", Desc: "CVX Book", Content: "Numerical algorithms for convex optimization problems"},
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello World")
	})
	router.HandleFunc("/articles/", func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode(Articles)
	}).Methods("GET")
	router.HandleFunc("/articles/{id}/", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		// id, _ := strconv.ParseInt(vars["id"], 10, 64)
		id := vars["id"]

		for _, article := range Articles {
			if article.Id == id {
				json.NewEncoder(rw).Encode(article)
			}
		}

	})
	router.HandleFunc("/articles/", func(rw http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		
		var newArticle Article

		json.Unmarshal(reqBody, &newArticle)

		Articles = append(Articles, newArticle)
		json.NewEncoder(rw).Encode(newArticle)

	}).Methods("POST")

	fmt.Println("Listening on port 8080 ...")
	http.ListenAndServe("0.0.0.0:8080", router)
}
