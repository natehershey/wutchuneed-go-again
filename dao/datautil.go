package dao

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

type Dao struct {
	Db            *sql.DB
	Driver        string
	ConnectString string
}

type WutchuneedDao interface {
	GetLists() ([]List, error)
	Init() Dao
}

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

// func Init() {
// 	db := getDbConnection()
// 	return db
// }

func (dao Dao) Init() Dao {
	log.Println("dao.Init()")

	dao.Driver = "mysql"
	dao.ConnectString = "root@tcp(localhost:3306)/wutchuneed?parseTime=true"
	db, err := dao.getDbConnection()
	if err != nil {
		log.Print("Failed in getDbConnection()")
		log.Print(err.Error())
		panic(err)
	}
	dao.Db = db
	return dao
}

func (dao Dao) getDbConnection() (*sql.DB, error) {
	log.Print("getDbConnection()")
	log.Print(dao.Driver)
	log.Print(dao.ConnectString)
	db, err := sql.Open(dao.Driver, dao.ConnectString)

	if err != nil {
		log.Print("first", err)
		panic(err)
	}
	log.Print("Pinging the DB")
	if err := db.Ping(); err != nil {
		log.Print(err)
		panic(err)
	}
	dao.Db = db
	return db, nil
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
