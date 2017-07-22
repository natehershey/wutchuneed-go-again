package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/natehershey/wutchuneed-go-again/dao"
	"github.com/natehershey/wutchuneed-go-again/routehandlers"
	"net/http"
)

func main() {
	fmt.Println("dao.Init()")
	dao.Init()
	r := mux.NewRouter()
	r.HandleFunc("/", routehandlers.HomeHandler)
	r.HandleFunc("/lists", routehandlers.GetListsHandler).Methods("GET")
	r.HandleFunc("/lists/{id:[0-9]+}", routehandlers.GetListHandler).Methods("GET")
	r.HandleFunc("/lists", routehandlers.PostListHandler).Methods("POST")

	r.HandleFunc("/categories", routehandlers.CategoriesHandler).Methods("GET")
	r.HandleFunc("/categories/{id:[0-9]+}", routehandlers.CategoryHandler).Methods("GET")

	r.HandleFunc("/items", routehandlers.ItemsHandler).Methods("GET")
	r.HandleFunc("/items/{id:[0-9]+}", routehandlers.ItemHandler).Methods("GET")

	http.Handle("/", r)
	fmt.Println("http.ListenAndServe('localhost:8081', nil)")
	fmt.Println(http.ListenAndServe("localhost:8081", nil))
}
