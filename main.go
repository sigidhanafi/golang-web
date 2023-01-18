package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	log.Println("Start listening port :8080")

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("Assets"))))
	mux.HandleFunc("/", handler.HomeHanlder)
	mux.HandleFunc("/products", handler.ProductListHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
