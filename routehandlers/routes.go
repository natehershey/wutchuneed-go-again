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
		writeErrorResponse(err.Error(), http.StatusNotFound, w, enc)
		return
	}

	json, err := json.Marshal(allLists)
	if err != nil {
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
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
		writeErrorResponse(err.Error(), http.StatusBadRequest, w, enc)
		return
	}

	list, err := dao.GetList(id)
	if err == sql.ErrNoRows {
		writeErrorResponse("List not found", http.StatusNotFound, w, enc)
		return
	}
	if err != nil {
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
		return
	}

	json, err := json.Marshal(list)
	if err != nil {
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
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
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
		return
	}

	json, err := json.Marshal(list)
	if err != nil {
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
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
		writeErrorResponse(err.Error(), http.StatusBadRequest, w, enc)
		return
	}

	success, err := dao.DeleteList(id)
	if err != nil {
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
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
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
		return
	}

	json, err := json.Marshal(allCategories)
	if err != nil {
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
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
		writeErrorResponse(err.Error(), http.StatusBadRequest, w, enc)
		return
	}

	category, err := dao.GetCategory(id)
	if err == sql.ErrNoRows {
		writeErrorResponse("Category not found", http.StatusNotFound, w, enc)
		return
	}
	if err != nil {
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
		return
	}

	json, err := json.Marshal(category)
	if err != nil {
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
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
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
		return
	}

	json, err := json.Marshal(category)
	if err != nil {
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
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
		writeErrorResponse(err.Error(), http.StatusBadRequest, w, enc)
		return
	}

	success, err := dao.DeleteCategory(id)
	if err != nil {
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
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
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
		return
	}

	json, err := json.Marshal(allItems)
	if err != nil {
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
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
		writeErrorResponse(err.Error(), http.StatusBadRequest, w, enc)
		return
	}

	item, err := dao.GetItem(id)
	if err == sql.ErrNoRows {
		writeErrorResponse("Item not found", http.StatusNotFound, w, enc)
		return
	}
	if err != nil {
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
		return
	}

	json, err := json.Marshal(item)
	if err != nil {
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
	}

	w.Write(json)
}

func PostItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	body, _ := ioutil.ReadAll(r.Body)
	item, err := dao.CreateItem(body)
	if err != nil {
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
		return
	}

	json, err := json.Marshal(item)
	if err != nil {
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
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
		writeErrorResponse(err.Error(), http.StatusBadRequest, w, enc)
		return
	}

	success, err := dao.DeleteItem(id)
	if err != nil {
		writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
		return
	}
	response := map[string]bool{"success": success}
	enc.Encode(response)
}

func writeErrorResponse(msg string, code int, w http.ResponseWriter, enc *json.Encoder) {
	w.WriteHeader(code)
	response := map[string]string{"error": msg}
	enc.Encode(response)
}
