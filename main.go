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
	r.HandleFunc("/status", routehandlers.StatusHandler)
	r.HandleFunc("/lists", routehandlers.GetListsHandler).Methods("GET")
	r.HandleFunc("/lists/{id:[0-9]+}", routehandlers.GetListHandler).Methods("GET")
	r.HandleFunc("/lists", routehandlers.PostListHandler).Methods("POST")
	r.HandleFunc("/lists/{id:[0-9]+}", routehandlers.DeleteListHandler).Methods("DELETE")

	r.HandleFunc("/categories", routehandlers.GetCategoriesHandler).Methods("GET")
	r.HandleFunc("/categories/{id:[0-9]+}", routehandlers.GetCategoryHandler).Methods("GET")
	r.HandleFunc("/categories", routehandlers.PostCategoryHandler).Methods("POST")
	r.HandleFunc("/categories/{id:[0-9]+}", routehandlers.DeleteCategoryHandler).Methods("DELETE")

	r.HandleFunc("/items", routehandlers.GetItemsHandler).Methods("GET")
	r.HandleFunc("/items/{id:[0-9]+}", routehandlers.GetItemHandler).Methods("GET")
	r.HandleFunc("/items", routehandlers.PostItemHandler).Methods("POST")
	r.HandleFunc("/items/{id:[0-9]+}", routehandlers.DeleteItemHandler).Methods("DELETE")

	http.Handle("/", r)

	fmt.Println("http.ListenAndServe('localhost:8081', nil)")
	fmt.Println(http.ListenAndServe("localhost:8081", nil))
}
