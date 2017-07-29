package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/natehershey/wutchuneed-go-again/dao"
	"github.com/natehershey/wutchuneed-go-again/routehandlers"
	"net/http"
	"strings"
)

func main() {
	// We'll use DI to pass the DB into the routes/handlers to avoid
	//  global state and make mocking in tests easier
	db := dao.Dao{}.Init()
	fmt.Printf("DB URL: %s\n", db.ConnectString)
	fmt.Printf("DB Db: %v\n", db.Db)

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	// r.HandleFunc("/", routehandlers.HomeHandler)
	r.HandleFunc("/status", routehandlers.StatusHandler)

	r.Path("/api/v1/lists").
		HandlerFunc(routehandlers.GetListsHandler(db)).
		Name("lists").
		Methods("GET")
	r.HandleFunc("/api/v1/lists/{id:[0-9]+}", routehandlers.GetListHandler(db)).Methods("GET")
	r.HandleFunc("/api/v1/lists", routehandlers.PostListHandler(db)).Methods("POST")
	r.HandleFunc("/api/v1/lists/{id:[0-9]+}", routehandlers.DeleteListHandler(db)).Methods("DELETE")

	r.HandleFunc("/api/v1/categories", routehandlers.GetCategoriesHandler(db)).Methods("GET")
	r.HandleFunc("/api/v1/categories/{id:[0-9]+}", routehandlers.GetCategoryHandler(db)).Methods("GET")
	r.HandleFunc("/api/v1/categories", routehandlers.PostCategoryHandler(db)).Methods("POST")
	r.HandleFunc("/api/v1/categories/{id:[0-9]+}", routehandlers.PutCategoryHandler(db)).Methods("PUT")
	r.HandleFunc("/api/v1/categories/{id:[0-9]+}", routehandlers.DeleteCategoryHandler(db)).Methods("DELETE")

	r.HandleFunc("/api/v1/items", routehandlers.GetItemsHandler(db)).Methods("GET")
	r.HandleFunc("/api/v1/items/{id:[0-9]+}", routehandlers.GetItemHandler(db)).Methods("GET")
	r.HandleFunc("/api/v1/items", routehandlers.PostItemHandler(db)).Methods("POST")
	r.HandleFunc("/api/v1/items/{id:[0-9]+}", routehandlers.PutItemHandler(db)).Methods("PUT")
	r.HandleFunc("/api/v1/items/{id:[0-9]+}", routehandlers.DeleteItemHandler(db)).Methods("DELETE")

	http.Handle("/", Middleware(r))

	fmt.Println("http.ListenAndServe('localhost:8081', nil)")
	fmt.Println(http.ListenAndServe("localhost:8081", nil))
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("\n", formatRequest(r), "\n")

		h.ServeHTTP(w, r)
	})
}

func formatRequest(r *http.Request) string {
	var request []string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	return strings.Join(request, "\n")
}
