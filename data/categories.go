package data

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetCategories() []Category {
	fmt.Printf("getCategories()\n")
	rows, err := db.Query("select id, name, list_id from categories;")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	foundCategories = []Category{}

	for rows.Next() {
		var c Category

		err = rows.Scan(&c.Id, &c.CategoryName, &c.ListId)
		if err != nil {
			log.Fatal(err)
		}

		c.Items = getItemsForCategory(c)

		foundCategories = append(foundCategories, c)
	}
	return foundCategories
}

func getCategoriesForList(l List) []Category {
	fmt.Printf("getCategoriesForList(l List)\n")
	rows, err := db.Query("select id, name from categories where list_id = ?;", l.Id)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	foundCategories = []Category{}

	for rows.Next() {
		var c Category

		err = rows.Scan(&c.Id, &c.CategoryName)
		if err != nil {
			log.Fatal(err)
		}

		c.Items = getItemsForCategoryAndList(l, c)

		foundCategories = append(foundCategories, c)
	}
	return foundCategories
}
