package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetCategories() ([]Category, error) {
	fmt.Printf("getCategories()\n")
	rows, err := db.Query("select id, name, list_id from categories;")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var foundCategories []Category

	for rows.Next() {
		var c Category

		err = rows.Scan(&c.Id, &c.CategoryName, &c.ListId)
		if err != nil {
			log.Fatal(err)
		}

		c.Items, err = getItemsForCategory(c)

		foundCategories = append(foundCategories, c)
	}
	return foundCategories, err
}

func GetCategory(id int) (Category, error) {
	fmt.Printf("getCategory(%d)\n", id)

	var c Category
	err := db.QueryRow("SELECT id, name, list_id FROM categories WHERE id = ?;", id).Scan(&c.Id, &c.CategoryName, &c.ListId)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No category with id %d.", id)
	case err != nil:
		log.Fatal(err)
	}

	c.Items, err = getItemsForCategory(c)
	return c, err
}

func getCategoriesForList(l List) ([]Category, error) {
	fmt.Printf("getCategoriesForList(l List)\n")
	rows, err := db.Query("select id, name from categories where list_id = ?;", l.Id)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var foundCategories []Category

	for rows.Next() {
		var c Category

		err = rows.Scan(&c.Id, &c.CategoryName)
		if err != nil {
			log.Fatal(err)
		}

		c.Items, err = getItemsForCategoryAndList(l, c)

		foundCategories = append(foundCategories, c)
	}
	return foundCategories, err
}