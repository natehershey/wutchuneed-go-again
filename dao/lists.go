package dao

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func (db Dao) GetLists() ([]List, error) {
	rows, err := db.Db.Query("select id, name, list_type, created_at, updated_at from lists;")
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

		l.Categories, err = db.GetCategoriesForList(l)
		if err != nil {
			log.Print(err)
		}

		foundLists = append(foundLists, l)
	}

	return foundLists, err
}

// func GetList(id int) (List, error) {
// 	var list List
// 	err := db.QueryRow("SELECT id, name, list_type, created_at, updated_at FROM lists WHERE id=?;", id).Scan(&list.Id, &list.Name, &list.Type, &list.CreatedAt, &list.UpdatedAt)

// 	if err != nil {
// 		return List{}, err
// 	}

// 	list.Categories, err = getCategoriesForList(list)
// 	if err != nil {
// 		return List{}, err
// 	}

// 	return list, err
// }

func (db Dao) GetList(id int) (List, error) {
	var list List
	err := db.Db.QueryRow("SELECT id, name, list_type, created_at, updated_at FROM lists WHERE id=?;", id).Scan(&list.Id, &list.Name, &list.Type, &list.CreatedAt, &list.UpdatedAt)

	if err != nil {
		return List{}, err
	}

	list.Categories, err = db.GetCategoriesForList(list)
	if err != nil {
		return List{}, err
	}

	return list, err
}

func (db Dao) CreateList(body []byte) (List, error) {
	var list List
	unmarshalErr := json.Unmarshal(body, &list)
	if unmarshalErr != nil {
		log.Print(unmarshalErr)
		return List{}, unmarshalErr
	}

	stmt, prepErr := db.Db.Prepare("INSERT INTO lists(name, list_type, created_at) VALUES(?,?, NOW())")
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
	queryErr := db.Db.QueryRow("SELECT id, name, list_type, created_at, updated_at FROM lists WHERE id=?;", lastId).Scan(&list.Id, &list.Name, &list.Type, &list.CreatedAt, &list.UpdatedAt)
	if queryErr != nil {
		log.Print(queryErr)
		return List{}, queryErr
	}
	return list, nil
}

func (db Dao) DeleteList(id int) (bool, error) {
	if id <= 0 {
		return false, fmt.Errorf("Bad list ID: %d", id)
	}

	stmt, err := db.Db.Prepare("DELETE FROM lists WHERE id = ?;")
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
