package routehandlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/natehershey/wutchuneed-go-again/dao"
	"io/ioutil"
	"net/http"
	"strconv"
)

//
// Route Handlers
//

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	response := map[string]bool{"alive": true}
	enc.Encode(response)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wutchuneed")
}

func GetListsHandler(db dao.Dao) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)

		allLists, err := db.GetLists()
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
	})
}

func GetListHandler(db dao.Dao) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			writeErrorResponse(err.Error(), http.StatusBadRequest, w, enc)
			return
		}

		list, err := db.GetList(id)
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
	})
}

func PostListHandler(db dao.Dao) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)

		body, _ := ioutil.ReadAll(r.Body)
		list, err := db.CreateList(body)
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
	})
}

func DeleteListHandler(db dao.Dao) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			writeErrorResponse(err.Error(), http.StatusBadRequest, w, enc)
			return
		}

		success, err := db.DeleteList(id)
		if err != nil {
			writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
			return
		}

		response := map[string]bool{"success": success}
		enc.Encode(response)
	})
}

func GetCategoriesHandler(db dao.Dao) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)

		allCategories, err := db.GetCategories()
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
	})
}

func GetCategoryHandler(db dao.Dao) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)

		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			writeErrorResponse(err.Error(), http.StatusBadRequest, w, enc)
			return
		}

		category, err := db.GetCategory(id)
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
	})
}

func PostCategoryHandler(db dao.Dao) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)

		body, _ := ioutil.ReadAll(r.Body)
		category, err := db.CreateCategory(body)
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
	})
}

func DeleteCategoryHandler(db dao.Dao) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			writeErrorResponse(err.Error(), http.StatusBadRequest, w, enc)
			return
		}

		success, err := db.DeleteCategory(id)
		if err != nil {
			writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
			return
		}

		response := map[string]bool{"success": success}
		enc.Encode(response)
	})
}

func GetItemsHandler(db dao.Dao) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)

		allItems, err := db.GetItems()
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
	})
}

func GetItemHandler(db dao.Dao) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			writeErrorResponse(err.Error(), http.StatusBadRequest, w, enc)
			return
		}

		item, err := db.GetItem(id)
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
	})
}

func PostItemHandler(db dao.Dao) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)

		body, _ := ioutil.ReadAll(r.Body)
		item, err := db.CreateItem(body)
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
	})
}

func DeleteItemHandler(db dao.Dao) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			writeErrorResponse(err.Error(), http.StatusBadRequest, w, enc)
			return
		}

		success, err := db.DeleteItem(id)
		if err != nil {
			writeErrorResponse(err.Error(), http.StatusInternalServerError, w, enc)
			return
		}
		response := map[string]bool{"success": success}
		enc.Encode(response)
	})
}

func writeErrorResponse(msg string, code int, w http.ResponseWriter, enc *json.Encoder) {
	w.WriteHeader(code)
	response := map[string]string{"error": msg}
	enc.Encode(response)
}
