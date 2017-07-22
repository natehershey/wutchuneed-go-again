package routehandlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/natehershey/wutchuneed-go-again/dao"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//
// Route Handlers
//

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get lists\n")
	allLists, err := dao.GetLists()
	fmt.Printf("Found: %+v\n", allLists)

	json, err := json.Marshal(allLists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func GetListsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get lists\n")
	allLists, err := dao.GetLists()
	if err != nil {
		log.Print(err)
	} else {
		fmt.Printf("Found: %+v\n", allLists)

		json, err := json.Marshal(allLists)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}

func GetListHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, idErr := strconv.Atoi(vars["id"])

	if idErr != nil {
		http.Error(w, idErr.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Get list %d\n", id)

	list, err := dao.GetList(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Printf("Found: %+v\n", list)

		json, err := json.Marshal(list)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}

func PostListHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	fmt.Printf("Creating list %v\n", body)
	list, err := dao.CreateList(body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {

		json, err := json.Marshal(list)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get categories\n")
	allCategories, err := dao.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Printf("Found: %+v\n", allCategories)

		json, err := json.Marshal(allCategories)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, idErr := strconv.Atoi(vars["id"])

	if idErr != nil {
		http.Error(w, idErr.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Geting category %d\n", id)
	category, err := dao.GetCategory(id)
	fmt.Printf("Found: %+v\n", category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		json, err := json.Marshal(category)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}

func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get items\n")
	allItems, err := dao.GetItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Printf("Found: %+v\n", allItems)

		json, err := json.Marshal(allItems)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}

func ItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, idErr := strconv.Atoi(vars["id"])

	if idErr != nil {
		http.Error(w, idErr.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Get item(%d)\n", id)
	item, err := dao.GetItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Printf("Found: %+v\n", item)

		json, err := json.Marshal(item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}
