package main

import "fmt"
import "net/http"
import "github.com/gorilla/mux"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/lists", ListsHandler)
	r.HandleFunc("/categories", CategoriesHandler)
	r.HandleFunc("/items", ItemsHandler)
	http.Handle("/", r)

	fmt.Println(http.ListenAndServe(":8081", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func ListsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Lists")
}

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Categories")
}

func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Items")
}
