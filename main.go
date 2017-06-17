package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/natehershey/wutchuneed-go-again/data"
	"github.com/natehershey/wutchuneed-go-again/routehandlers"
	"net/http"
)

func main() {
	data.Init()
	r := mux.NewRouter()
	r.HandleFunc("/", routehandlers.HomeHandler)
	r.HandleFunc("/lists", routehandlers.ListsHandler)
	r.HandleFunc("/categories", routehandlers.CategoriesHandler)
	r.HandleFunc("/items", routehandlers.ItemsHandler)
	http.Handle("/", r)
	fmt.Println(http.ListenAndServe("localhost:8081", nil))
}
