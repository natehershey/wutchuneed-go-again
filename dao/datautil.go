package dao

import (
	"database/sql"
	// "encoding/json"
	"bytes"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

var err error
var db *sql.DB

// var foundLists = []List{}
// var foundCategories = []Category{}
// var foundItems = []Item{}

var listDeleteStmt *sql.Stmt

type List struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	CreatedAt  NullTime   `json:"createdAt"`
	UpdatedAt  NullTime   `json:"updatedAt"`
	Categories []Category `json:"categories"`
}

type Category struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	ListId    int      `json:"listId"`
	CreatedAt NullTime `json:"createdAt"`
	UpdatedAt NullTime `json:"updatedAt"`
	Items     []Item   `json:"items"`
}

type Item struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Quantity   int        `json:"quantity"`
	Measure    NullString `json:"measure"`
	Status     NullString `json:"status"`
	CategoryId int        `json:"categoryId"`
	ListId     int        `json:"listId"`
	CreatedAt  NullTime   `json:"createdAt"`
	UpdatedAt  NullTime   `json:"updatedAt"`
}

type NullString struct {
	sql.NullString
}

type NullTime struct {
	mysql.NullTime
}

func Init() {
	getDbConnection()
}

func getDbConnection() {
	fmt.Println("Open DB Connection")
	db, err = sql.Open("mysql", "root@tcp(localhost:3306)/wutchuneed?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	var str = strings.Trim(string(data), `"`)
	if str == "null" {
		s.Valid = false
	} else {
		s.String = strings.Trim(string(data), `"`)
		s.Valid = true
	}
	return nil
}

func (s NullString) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer
	// buffer := bytes.NewBufferString("\"time\":")
	if !s.Valid {
		buffer.WriteString("null")
	} else {
		fmt.Fprintf(&buffer, "\"%s\"", s.String)
	}
	b := buffer.Bytes()
	log.Print(b)
	return b, nil
}

func (t NullTime) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer
	// buffer := bytes.NewBufferString("\"time\":")
	if !t.Valid {
		buffer.WriteString("null")
	} else {
		fmt.Fprintf(&buffer, "\"%s\"", t.String())
	}
	b := buffer.Bytes()
	log.Print(b)
	return b, nil
}

func (t NullTime) String() string {
	return fmt.Sprintf(t.Time.String())
}
