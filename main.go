package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	// "os"
)

var db *sql.DB
var err error
var foundLists = []List{}
var foundCategories = []Category{}
var foundItems = []Item{}

type List struct {
	Id         int        `json:"id"`
	ListName   string     `json:"name"`
	ListType   string     `json:"listType"`
	Categories []Category `json:"categories"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"name"`
	ListId       int    `json:"listId"`
	Items        []Item `json:"items"`
}

type Item struct {
	Id         int            `json:"id"`
	ItemName   string         `json:"name"`
	Quantity   int            `json:"quantity"`
	Measure    sql.NullString `json:"measure"`
	Status     sql.NullString `json:"status"`
	CategoryId int            `json:"categoryId"`
	ListId     int            `json:"listId"`
}

func main() {
	getDbConnection()

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/lists", ListsHandler)
	http.Handle("/", r)
	fmt.Println(http.ListenAndServe("localhost:8081", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get lists\n")
	allLists := getLists()
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
	allLists := getLists()
	fmt.Printf("Found: %+v\n", allLists)

	json, err := json.Marshal(allLists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func getLists() []List {
	fmt.Printf("'select id, name, list_type from lists;' in getLists()\n")
	rows, err := db.Query("select id, name, list_type from lists;")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Printf("select finished\n")

	foundLists = []List{}

	for rows.Next() {
		var l List

		fmt.Printf("Scanning rows\n")
		err = rows.Scan(&l.Id, &l.ListName, &l.ListType)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Finished scanning rows\n")

		l.Categories = getCategoriesForList(l)

		foundLists = append(foundLists, l)
	}

	return foundLists
}

func getCategoriesForList(l List) []Category {
	fmt.Printf("getCategoriesForList(l List)\n")
	rows, err := db.Query("select id, name from categories where list_id = ?;", l.Id)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	foundCategories = []Category{}

	for rows.Next() {
		var c Category

		err = rows.Scan(&c.Id, &c.CategoryName)
		if err != nil {
			log.Fatal(err)
		}

		c.Items = getItemsForCategory(l, c)

		foundCategories = append(foundCategories, c)
	}
	return foundCategories
}

func getItemsForCategory(l List, c Category) []Item {
	log.Print("select id, name, quantity, measure, status from items where category_id = %s;", c.Id)
	rows, err := db.Query("SELECT id, name, quantity, measure, status FROM items WHERE list_id = ? AND category_id = ?;", l.Id, c.Id)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Printf("got rows\n")
	foundItems = []Item{}

	for rows.Next() {
		var i Item

		fmt.Printf("Scanning into item\n")
		err = rows.Scan(&i.Id, &i.ItemName, &i.Quantity, &i.Measure, &i.Status)
		if err != nil {
			log.Fatal(err)
		}

		foundItems = append(foundItems, i)
		fmt.Printf("Name: %s, ID: %d, status: %s", i.ItemName, i.Id, i.Status)
	}
	return foundItems
}

func getDbConnection() {
	db, err = sql.Open("mysql", "root@tcp(localhost:3306)/wutchuneed")

	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}
