package data

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetItems() []Item {
	log.Print("SELECT id, name, quantity, measure, status FROM items;")
	rows, err := db.Query("SELECT id, name, quantity, measure, status FROM items;")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Printf("got rows\n")
	foundItems = []Item{}

	for rows.Next() {
		var i Item

		fmt.Printf("Scanning into item\n")
		err = rows.Scan(&i.Id, &i.ItemName, &i.Quantity, &i.Measure, &i.Status)
		if err != nil {
			log.Fatal(err)
		}

		foundItems = append(foundItems, i)
		fmt.Printf("Name: %s, ID: %d, status: %s", i.ItemName, i.Id, i.Status)
	}
	return foundItems
}

func getItemsForCategoryAndList(l List, c Category) []Item {
	log.Print("select id, name, quantity, measure, status from items where category_id = %s;", c.Id)
	rows, err := db.Query("SELECT id, name, quantity, measure, status FROM items WHERE list_id = ? AND category_id = ?;", l.Id, c.Id)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Printf("got rows\n")
	foundItems = []Item{}

	for rows.Next() {
		var i Item

		fmt.Printf("Scanning into item\n")
		err = rows.Scan(&i.Id, &i.ItemName, &i.Quantity, &i.Measure, &i.Status)
		if err != nil {
			log.Fatal(err)
		}

		foundItems = append(foundItems, i)
		fmt.Printf("Name: %s, ID: %d, status: %s", i.ItemName, i.Id, i.Status)
	}
	return foundItems
}

func getItemsForCategory(c Category) []Item {
	log.Print("select id, name, quantity, measure, status from items where category_id = %s;", c.Id)
	rows, err := db.Query("SELECT id, name, quantity, measure, status FROM items WHERE category_id = ?;", c.Id)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Printf("got rows\n")
	foundItems = []Item{}

	for rows.Next() {
		var i Item

		fmt.Printf("Scanning into item\n")
		err = rows.Scan(&i.Id, &i.ItemName, &i.Quantity, &i.Measure, &i.Status)
		if err != nil {
			log.Fatal(err)
		}

		foundItems = append(foundItems, i)
		fmt.Printf("Name: %s, ID: %d, status: %s", i.ItemName, i.Id, i.Status)
	}
	return foundItems
}
