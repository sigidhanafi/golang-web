package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func HomeHanlder(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)

		return
	}

	tmplt, err := template.ParseFiles(path.Join("views", "home.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Server error, failed to load view", http.StatusInternalServerError)
		return
	}

	data := map[string]string{"title": "Learn Golang", "content": "Hello, welcome, we learn golang enjoy it!"}

	err = tmplt.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server error, try again", http.StatusInternalServerError)
		return
	}
}

func ProductListHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/products" {
		http.NotFound(w, r)
		return
	}

	tmplt, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Server error, try again", http.StatusInternalServerError)
		return
	}

	products := []string{
		"Product One",
		"Product Two",
		"product Three",
	}
	data := map[string]interface{}{"title": "Learn Golang", "content": products}

	err = tmplt.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server error, try again", http.StatusInternalServerError)
		return
	}
}
