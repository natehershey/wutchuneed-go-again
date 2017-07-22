package dao

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetLists() ([]List, error) {
	fmt.Printf("'select id, name, list_type from lists;' in GetLists()\n")

	var foundLists []List

	rows, err := db.Query("select id, name, list_type from lists;")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Printf("select finished\n")

	for rows.Next() {
		var l List

		fmt.Printf("Scanning rows\n")
		err = rows.Scan(&l.Id, &l.ListName, &l.ListType)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Finished scanning rows\n")

		l.Categories, err = getCategoriesForList(l)

		foundLists = append(foundLists, l)
	}

	return foundLists, err
}

func GetList(id int) (List, error) {
	fmt.Printf("'select id, name, list_type from lists where id = %d;' in GetLists()\n", id)

	var list List
	err := db.QueryRow("SELECT id, name, list_type FROM lists WHERE id=?;", id).Scan(&list.Id, &list.ListName, &list.ListType)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No list with id %d.", id)
	case err != nil:
		log.Fatal(err)
	}

	list.Categories, err = getCategoriesForList(list)

	return list, err
}

func CreateList(body []byte) (List, error) {
	fmt.Println("Inserting list %v\n", body)

	var list List
	unmarshalErr := json.Unmarshal(body, &list)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}

	stmt, prepErr := db.Prepare("INSERT INTO lists(name, list_type, created_at) VALUES(?,?, NOW())")
	if prepErr != nil {
		log.Fatal(prepErr)
	}
	res, insertErr := stmt.Exec(list.ListName, list.ListType)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	lastId, lastIdErr := res.LastInsertId()
	if lastIdErr != nil {
		log.Fatal(lastIdErr)
	}
	queryErr := db.QueryRow("SELECT id, name, list_type FROM lists WHERE id=?;", lastId).Scan(&list.Id, &list.ListName, &list.ListType)
	if queryErr != nil {
		log.Fatal(queryErr)
	}
	return list, err
}

func DeleteList(body []byte) (bool, error) {
	return false, err
}
