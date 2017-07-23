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
	GetListsHandler(w, r)
}

func GetListsHandler(w http.ResponseWriter, r *http.Request) {
	allLists, err := dao.GetLists()
	if err != nil {
		log.Print(err)
	} else {
		json, err := json.Marshal(allLists)
		if err != nil {
			log.Print(err)
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

	list, err := dao.GetList(id)
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

func DeleteListHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	success, err := dao.DeleteList(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)

		response := map[string]bool{"success": success}
		enc.Encode(response)
	}
}

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
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

func GetCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, idErr := strconv.Atoi(vars["id"])

	if idErr != nil {
		http.Error(w, idErr.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Getting category %d\n", id)
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

func PostCategoryHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	fmt.Printf("Creating category %v\n", body)
	category, err := dao.CreateCategory(body)

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

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	success, err := dao.DeleteCategory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)

		response := map[string]bool{"success": success}
		enc.Encode(response)
	}
}

func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
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

func GetItemHandler(w http.ResponseWriter, r *http.Request) {
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

func PostItemHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	fmt.Printf("Creating item %v\n", body)
	item, err := dao.CreateItem(body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {

		json, err := json.Marshal(item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	success, err := dao.DeleteItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)

		response := map[string]bool{"success": success}
		enc.Encode(response)
	}
}
