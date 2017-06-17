package data

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetLists() []List {
	fmt.Printf("'select id, name, list_type from lists;' in GetLists()\n")
	rows, err := db.Query("select id, name, list_type from lists;")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Printf("select finished\n")

	foundLists = []List{}

	for rows.Next() {
		var l List

		fmt.Printf("Scanning rows\n")
		err = rows.Scan(&l.Id, &l.ListName, &l.ListType)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Finished scanning rows\n")

		l.Categories = getCategoriesForList(l)

		foundLists = append(foundLists, l)
	}

	return foundLists
}
