package routehandlers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/natehershey/wutchuneed-go-again/dao"
	"io/ioutil"
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
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	allLists, err := dao.GetLists()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	json, err := json.Marshal(allLists)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	w.Write(json)
}

func GetListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	list, err := dao.GetList(id)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{"error": "List not found"}
		enc.Encode(response)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	json, err := json.Marshal(list)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	w.Write(json)
}

func PostListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	body, _ := ioutil.ReadAll(r.Body)
	list, err := dao.CreateList(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	json, err := json.Marshal(list)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	w.Write(json)
}

func DeleteListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	success, err := dao.DeleteList(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	response := map[string]bool{"success": success}
	enc.Encode(response)
}

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	allCategories, err := dao.GetCategories()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	json, err := json.Marshal(allCategories)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	w.Write(json)
}

func GetCategoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	category, err := dao.GetCategory(id)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{"error": "Category not found"}
		enc.Encode(response)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	json, err := json.Marshal(category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	w.Write(json)
}

func PostCategoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	body, _ := ioutil.ReadAll(r.Body)
	category, err := dao.CreateCategory(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	json, err := json.Marshal(category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	w.Write(json)
}

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	success, err := dao.DeleteCategory(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	response := map[string]bool{"success": success}
	enc.Encode(response)
}

func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	allItems, err := dao.GetItems()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	json, err := json.Marshal(allItems)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	w.Write(json)
}

func GetItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	item, err := dao.GetItem(id)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{"error": "Item not found"}
		enc.Encode(response)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	json, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	w.Write(json)
}

func PostItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	body, _ := ioutil.ReadAll(r.Body)
	item, err := dao.CreateItem(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	json, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	w.Write(json)
}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}

	success, err := dao.DeleteItem(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		enc.Encode(response)
		return
	}
	response := map[string]bool{"success": success}
	enc.Encode(response)
}
