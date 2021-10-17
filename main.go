package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func main() {
	// database data
	Articles = []Article{
		{Id: 1, Title: "Model Predictive Control", Desc: "MPC Book", Content: "Model predictive control"},
		{Id: 2, Title: "Convex Optimization", Desc: "CVX Book", Content: "Numerical algorithms for convex optimization problems"},
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
		id, _ := strconv.ParseInt(vars["id"], 10, 64)

		for _, article := range Articles {
			if article.Id == id {
				json.NewEncoder(rw).Encode(article)
			}
		}

	}).Methods("GET")
	router.HandleFunc("/articles/", func(rw http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)

		var newArticle Article

		json.Unmarshal(reqBody, &newArticle)

		Articles = append(Articles, newArticle)
		json.NewEncoder(rw).Encode(newArticle)

	}).Methods("POST")

	router.HandleFunc("/articles/{id}/", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		for index, v := range Articles {
			if v.Id == id {
				Articles = append(Articles[:index], Articles[index + 1:]...)
			}
		}
	}).Methods("DELETE")
	
	router.HandleFunc("/articles/{id}/", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		
		reqBody, _ := ioutil.ReadAll(r.Body)

		var newArticle Article

		json.Unmarshal(reqBody, &newArticle)

		for index, v := range Articles {
			if v.Id == id {
				Articles[index] = newArticle
			}
		}
	}).Methods("PUT")
	fmt.Println("Listening on port 8080 ...")
	http.ListenAndServe("0.0.0.0:8080", router)
}
