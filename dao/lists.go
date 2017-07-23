package dao

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetLists() ([]List, error) {
	rows, err := db.Query("select id, name, list_type, created_at, updated_at from lists;")
	if err != nil {
		return []List{}, err
	}
	defer rows.Close()

	var foundLists = []List{}
	for rows.Next() {
		var l List
		err = rows.Scan(&l.Id, &l.Name, &l.Type, &l.CreatedAt, &l.UpdatedAt)
		if err != nil {
			log.Print(err)
		}

		l.Categories, err = getCategoriesForList(l)
		if err != nil {
			log.Print(err)
		}

		foundLists = append(foundLists, l)
	}

	return foundLists, err
}

func GetList(id int) (List, error) {
	var list List
	err := db.QueryRow("SELECT id, name, list_type, created_at, updated_at FROM lists WHERE id=?;", id).Scan(&list.Id, &list.Name, &list.Type, &list.CreatedAt, &list.UpdatedAt)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No list with id %d.", id)
	case err != nil:
		return List{}, err
	}

	list.Categories, err = getCategoriesForList(list)
	if err != nil {
		return List{}, err
	}

	return list, err
}

func CreateList(body []byte) (List, error) {
	var list List
	unmarshalErr := json.Unmarshal(body, &list)
	if unmarshalErr != nil {
		log.Print(unmarshalErr)
		return List{}, unmarshalErr
	}

	stmt, prepErr := db.Prepare("INSERT INTO lists(name, list_type, created_at) VALUES(?,?, NOW())")
	if prepErr != nil {
		log.Print(prepErr)
		return List{}, prepErr
	}
	res, insertErr := stmt.Exec(list.Name, list.Type)
	if insertErr != nil {
		log.Print(insertErr)
		return List{}, insertErr
	}
	lastId, lastIdErr := res.LastInsertId()
	if lastIdErr != nil {
		log.Print(lastIdErr)
		return List{}, lastIdErr
	}
	queryErr := db.QueryRow("SELECT id, name, list_type, created_at, updated_at FROM lists WHERE id=?;", lastId).Scan(&list.Id, &list.Name, &list.Type, &list.CreatedAt, &list.UpdatedAt)
	if queryErr != nil {
		log.Print(queryErr)
		return List{}, queryErr
	}
	return list, nil
}

func DeleteList(id int) (bool, error) {
	if id <= 0 {
		return false, fmt.Errorf("Bad list ID: %d", id)
	}

	stmt, err := db.Prepare("DELETE FROM lists WHERE id = ?;")
	if err != nil {
		log.Print(err)
		return false, err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		log.Print(err)
		return false, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Print(err)
		return false, err
	}
	if rowsAffected == 0 {
		return false, fmt.Errorf("Unable to delete list with ID %d", id)
	}
	return true, nil
}
