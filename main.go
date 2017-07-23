package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/natehershey/wutchuneed-go-again/dao"
	"github.com/natehershey/wutchuneed-go-again/routehandlers"
	"net/http"
)

func main() {
	// We'll use DI to pass the DB into the routes/handlers to avoid
	//  global state and make mocking in tests easier
	db := dao.Dao{}.Init()
	fmt.Printf("DB URL: %s\n", db.ConnectString)
	fmt.Printf("DB Db: %v\n", db.Db)

	r := mux.NewRouter()
	r.HandleFunc("/", routehandlers.HomeHandler)
	r.HandleFunc("/status", routehandlers.StatusHandler)
	r.Path("/lists").
		HandlerFunc(routehandlers.GetListsHandler(db)).
		Name("lists").
		Methods("GET")
	r.HandleFunc("/lists/{id:[0-9]+}", routehandlers.GetListHandler(db)).Methods("GET")
	r.HandleFunc("/lists", routehandlers.PostListHandler(db)).Methods("POST")
	r.HandleFunc("/lists/{id:[0-9]+}", routehandlers.DeleteListHandler(db)).Methods("DELETE")

	r.HandleFunc("/categories", routehandlers.GetCategoriesHandler(db)).Methods("GET")
	r.HandleFunc("/categories/{id:[0-9]+}", routehandlers.GetCategoryHandler(db)).Methods("GET")
	r.HandleFunc("/categories", routehandlers.PostCategoryHandler(db)).Methods("POST")
	r.HandleFunc("/categories/{id:[0-9]+}", routehandlers.DeleteCategoryHandler(db)).Methods("DELETE")

	r.HandleFunc("/items", routehandlers.GetItemsHandler(db)).Methods("GET")
	r.HandleFunc("/items/{id:[0-9]+}", routehandlers.GetItemHandler(db)).Methods("GET")
	r.HandleFunc("/items", routehandlers.PostItemHandler(db)).Methods("POST")
	r.HandleFunc("/items/{id:[0-9]+}", routehandlers.DeleteItemHandler(db)).Methods("DELETE")

	http.Handle("/", r)

	fmt.Println("http.ListenAndServe('localhost:8081', nil)")
	fmt.Println(http.ListenAndServe("localhost:8081", nil))
}
