package routehandlers

import (
	"encoding/json"
	"fmt"
	"github.com/natehershey/wutchuneed-go-again/data"
	"net/http"
)

//
// Route Handlers
//

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get lists\n")
	allLists := data.GetLists()
	fmt.Printf("Found: %+v\n", allLists)

	json, err := json.Marshal(allLists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func ListsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get lists\n")
	allLists := data.GetLists()
	fmt.Printf("Found: %+v\n", allLists)

	json, err := json.Marshal(allLists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get categories\n")
	allCategories := data.GetCategories()
	fmt.Printf("Found: %+v\n", allCategories)

	json, err := json.Marshal(allCategories)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get items\n")
	allItems := data.GetItems()
	fmt.Printf("Found: %+v\n", allItems)

	json, err := json.Marshal(allItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
