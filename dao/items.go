package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetItems() ([]Item, error) {
	rows, err := db.Query("SELECT id, name, quantity, measure, status, category_id, list_id FROM items;")
	if err != nil {
		log.Print(err)
		return []Item{}, err
	}
	defer rows.Close()

	var foundItems []Item

	for rows.Next() {
		var i Item

		err = rows.Scan(&i.Id, &i.ItemName, &i.Quantity, &i.Measure, &i.Status, &i.CategoryId, &i.ListId)
		if err != nil {
			log.Print(err)
		}

		foundItems = append(foundItems, i)
	}
	return foundItems, err
}

func GetItem(id int) (Item, error) {
	var i Item

	err := db.QueryRow("SELECT id, name, quantity, measure, status, category_id, list_id FROM items where id = ?;", id).Scan(&i.Id, &i.ItemName, &i.Quantity, &i.Measure, &i.Status, &i.CategoryId, &i.ListId)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No item with id %d.", id)
	case err != nil:
		log.Print(err)
		return Item{}, err
	}
	return i, err
}

func getItemsForCategoryAndList(l List, c Category) ([]Item, error) {
	rows, err := db.Query("SELECT id, name, quantity, measure, status, category_id, list_id FROM items WHERE list_id = ? AND category_id = ?;", l.Id, c.Id)
	if err != nil {
		log.Print(err)
		return []Item{}, err
	}
	defer rows.Close()

	var foundItems []Item
	for rows.Next() {
		var i Item

		err = rows.Scan(&i.Id, &i.ItemName, &i.Quantity, &i.Measure, &i.Status, &i.CategoryId, &i.ListId)
		if err != nil {
			log.Print(err)
		}

		foundItems = append(foundItems, i)
	}
	return foundItems, err
}

func getItemsForCategory(c Category) ([]Item, error) {
	rows, err := db.Query("SELECT id, name, quantity, measure, status, category_id, list_id FROM items WHERE category_id = ?;", c.Id)
	if err != nil {
		log.Print(err)
		return []Item{}, err
	}
	defer rows.Close()

	var foundItems []Item
	for rows.Next() {
		var i Item

		err = rows.Scan(&i.Id, &i.ItemName, &i.Quantity, &i.Measure, &i.Status, &i.CategoryId, &i.ListId)
		if err != nil {
			log.Print(err)
			return []Item{}, err
		}

		foundItems = append(foundItems, i)
	}
	return foundItems, err
}
