package dao

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetCategories() ([]Category, error) {
	rows, err := db.Query("select id, name, list_id, created_at, updated_at from categories;")

	if err != nil {
		log.Print(err)
		return []Category{}, err
	}
	defer rows.Close()

	var foundCategories = []Category{}

	for rows.Next() {
		var c Category

		err = rows.Scan(&c.Id, &c.Name, &c.ListId, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return []Category{}, err
		}

		c.Items, err = getItemsForCategory(c)
		if err != nil {
			log.Print(err)
			return []Category{}, err
		}

		foundCategories = append(foundCategories, c)
	}
	return foundCategories, err
}

func GetCategory(id int) (Category, error) {
	var c Category
	err := db.QueryRow("SELECT id, name, list_id, created_at, updated_at FROM categories WHERE id = ?;", id).Scan(&c.Id, &c.Name, &c.ListId, &c.CreatedAt, &c.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No category with id %d.", id)
		return Category{}, nil
	case err != nil:
		log.Print(err)
		return Category{}, nil
	}

	c.Items, err = getItemsForCategory(c)
	if err != nil {
		log.Print(err)
	}
	return c, err
}

func CreateCategory(body []byte) (Category, error) {
	var category Category
	unmarshalErr := json.Unmarshal(body, &category)
	if unmarshalErr != nil {
		log.Print(unmarshalErr)
		return Category{}, unmarshalErr
	}

	stmt, prepErr := db.Prepare("INSERT INTO categories(name, list_id, created_at) VALUES(?,?, NOW())")
	if prepErr != nil {
		log.Print(prepErr)
		return Category{}, prepErr
	}
	res, insertErr := stmt.Exec(category.Name, category.ListId)
	if insertErr != nil {
		log.Print(insertErr)
		return Category{}, insertErr
	}
	lastId, lastIdErr := res.LastInsertId()
	if lastIdErr != nil {
		log.Print(lastIdErr)
		return Category{}, lastIdErr
	}
	queryErr := db.QueryRow("SELECT id, name, list_id, created_at, updated_at FROM categories WHERE id=?;", lastId).Scan(&category.Id, &category.Name, &category.ListId, &category.CreatedAt, &category.UpdatedAt)
	if queryErr != nil {
		log.Print(queryErr)
		return Category{}, queryErr
	}
	return category, nil
}

func DeleteCategory(id int) (bool, error) {
	if id <= 0 {
		return false, fmt.Errorf("Bad category ID: %d", id)
	}

	stmt, err := db.Prepare("DELETE FROM categories WHERE id = ?;")
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
		return false, fmt.Errorf("Unable to delete category with ID %d", id)
	}
	return true, nil
}

func getCategoriesForList(l List) ([]Category, error) {
	rows, err := db.Query("select id, name, list_id, created_at, updated_at from categories where list_id = ?;", l.Id)
	if err != nil {
		log.Print(err)
		return []Category{}, err
	}
	defer rows.Close()

	var foundCategories = []Category{}

	for rows.Next() {
		var c Category

		err = rows.Scan(&c.Id, &c.Name, &c.ListId, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			log.Print(err)
		}

		c.Items, err = getItemsForCategoryAndList(l, c)
		if err != nil {
			log.Print(err)
		}

		foundCategories = append(foundCategories, c)
	}
	return foundCategories, err
}
