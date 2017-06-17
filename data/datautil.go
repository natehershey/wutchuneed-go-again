package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var err error
var db *sql.DB
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

func getDbConnection() {
	db, err = sql.Open("mysql", "root@tcp(localhost:3306)/wutchuneed")

	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func Init() {
	getDbConnection()
}
