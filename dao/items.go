package dao

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func (db Dao) GetItems() ([]Item, error) {
	rows, err := db.Db.Query("SELECT id, name, quantity, measure, status, category_id, list_id, created_at, updated_at FROM items;")
	if err != nil {
		log.Print(err)
		return []Item{}, err
	}
	defer rows.Close()

	var foundItems = []Item{}

	for rows.Next() {
		var i Item

		err = rows.Scan(&i.Id, &i.Name, &i.Quantity, &i.Measure, &i.Status, &i.CategoryId, &i.ListId, &i.CreatedAt, &i.UpdatedAt)
		if err != nil {
			log.Print(err)
		}

		foundItems = append(foundItems, i)
	}
	return foundItems, err
}

func (db Dao) GetItem(id int) (Item, error) {
	var i Item

	err := db.Db.QueryRow("SELECT id, name, quantity, measure, status, category_id, list_id FROM items where id = ?;", id).Scan(&i.Id, &i.Name, &i.Quantity, &i.Measure, &i.Status, &i.CategoryId, &i.ListId)
	if err != nil {
		log.Print(err)
		return Item{}, err
	}
	return i, err
}

func (db Dao) CreateItem(body []byte) (Item, error) {
	var item Item
	unmarshalErr := json.Unmarshal(body, &item)
	if unmarshalErr != nil {
		log.Print(unmarshalErr)
		return Item{}, unmarshalErr
	}

	stmt, prepErr := db.Db.Prepare("INSERT INTO items(name, quantity, measure, status, list_id, category_id, created_at) VALUES(?,?,?,?,?,?,NOW())")
	if prepErr != nil {
		log.Print(prepErr)
		return Item{}, prepErr
	}
	res, insertErr := stmt.Exec(item.Name, item.Quantity, item.Measure, item.Status, item.ListId, item.CategoryId)
	if insertErr != nil {
		log.Print(insertErr)
		return Item{}, insertErr
	}
	lastId, lastIdErr := res.LastInsertId()
	if lastIdErr != nil {
		log.Print(lastIdErr)
		return Item{}, lastIdErr
	}
	queryErr := db.Db.QueryRow("SELECT id, name, quantity, measure, status, list_id, category_id, created_at, updated_at FROM items WHERE id=?;", lastId).Scan(&item.Id, &item.Name, &item.Quantity, &item.Measure, &item.Status, &item.ListId, &item.CategoryId, &item.CreatedAt, &item.UpdatedAt)
	if queryErr != nil {
		log.Print(queryErr)
		return Item{}, queryErr
	}
	return item, nil
}

func (db Dao) UpdateItem(id int, body []byte) (Item, error) {
	var item Item
	unmarshalErr := json.Unmarshal(body, &item)
	log.Println("Unmarshal")
	log.Println(body)
	if unmarshalErr != nil {
		log.Print(unmarshalErr)
		return Item{}, unmarshalErr
	}
	log.Println("Body: ", body)
	log.Println("Item: ", item)

	var nameValueBuffer bytes.Buffer
	// TODO: Don't get pwned by SQL injection
	if item.Name != "" {
		s := fmt.Sprintf("%s = '%s' ", "name", item.Name)
		nameValueBuffer.WriteString(s)
	}
	log.Print(item.Quantity)

	if item.Quantity > 0 {
		s := fmt.Sprintf("%s = %s ", "quantity", item.Quantity)
		nameValueBuffer.WriteString(s)
	}
	if item.Measure.Valid {
		s := fmt.Sprintf("%s = '%s' ", "measure", item.Measure.String)
		nameValueBuffer.WriteString(s)
	}
	if item.Status.Valid {
		s := fmt.Sprintf("%s = '%s' ", "status", item.Status.String)
		nameValueBuffer.WriteString(s)
	}
	log.Print(item.ListId)
	if item.ListId > 0 {
		s := fmt.Sprintf("%s = %s ", "list_id", item.ListId)
		nameValueBuffer.WriteString(s)
	}
	log.Print(item.CategoryId)
	if item.CategoryId > 0 {
		s := fmt.Sprintf("%s = %s ", "category_id", item.CategoryId)
		nameValueBuffer.WriteString(s)
	}

	sql := fmt.Sprintf("UPDATE items SET %sWHERE id = %d;", nameValueBuffer.String(), id)
	log.Println(sql)
	res, updateErr := db.Db.Exec(sql)
	if updateErr != nil {
		log.Print(updateErr)
		return Item{}, updateErr
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Print(err)
		return Item{}, err
	}
	if rowsAffected == 0 {
		return Item{}, fmt.Errorf("Unable to update item with ID %d", id)
	}

	queryErr := db.Db.QueryRow("SELECT id, name, quantity, measure, status, list_id, category_id, created_at, updated_at FROM items WHERE id = ?;", id).Scan(&item.Id, &item.Name, &item.Quantity, &item.Measure, &item.Status, &item.ListId, &item.CategoryId, &item.CreatedAt, &item.UpdatedAt)
	if queryErr != nil {
		log.Print(queryErr)
		return Item{}, queryErr
	}
	return item, nil
}

func (db Dao) DeleteItem(id int) (bool, error) {
	if id <= 0 {
		return false, fmt.Errorf("Bad item ID: %d", id)
	}

	stmt, err := db.Db.Prepare("DELETE FROM items WHERE id = ?;")
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
		return false, fmt.Errorf("Unable to delete item with ID %d", id)
	}
	return true, nil
}

func (db Dao) getItemsForCategory(c Category) ([]Item, error) {
	rows, err := db.Db.Query("SELECT id, name, quantity, measure, status, category_id, list_id, created_at, updated_at FROM items WHERE category_id = ?;", c.Id)
	if err != nil {
		log.Print(err)
		return []Item{}, err
	}
	defer rows.Close()

	var foundItems = []Item{}
	for rows.Next() {
		var i Item

		err = rows.Scan(&i.Id, &i.Name, &i.Quantity, &i.Measure, &i.Status, &i.CategoryId, &i.ListId, &i.CreatedAt, &i.UpdatedAt)
		if err != nil {
			log.Print(err)
			return []Item{}, err
		}

		foundItems = append(foundItems, i)
	}
	printSlice(foundItems)
	return foundItems, err
}

func (db Dao) GetItemsForCategoryAndList(l List, c Category) ([]Item, error) {
	rows, err := db.Db.Query("SELECT id, name, quantity, measure, status, category_id, list_id, created_at, updated_at FROM items WHERE list_id = ? AND category_id = ?;", l.Id, c.Id)
	if err != nil {
		log.Print(err)
		return []Item{}, err
	}
	defer rows.Close()

	var foundItems = []Item{}
	for rows.Next() {
		var i Item

		err = rows.Scan(&i.Id, &i.Name, &i.Quantity, &i.Measure, &i.Status, &i.CategoryId, &i.ListId, &i.CreatedAt, &i.UpdatedAt)
		if err != nil {
			log.Print(err)
		}

		foundItems = append(foundItems, i)
	}
	log.Printf("foundItems:")
	printSlice(foundItems)
	return foundItems, err
}

func printSlice(s []Item) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
