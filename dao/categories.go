package dao

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type ValidationError struct {
	Message string
}

func (err ValidationError) Error() string {
	return fmt.Sprintf("Validation error: %s", err.Message)
}

func (db Dao) GetCategories() ([]Category, error) {
	rows, err := db.Db.Query("select id, name, list_id, created_at, updated_at from categories;")

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

		c.Items, err = db.getItemsForCategory(c)
		if err != nil {
			log.Print(err)
			return []Category{}, err
		}

		foundCategories = append(foundCategories, c)
	}
	return foundCategories, err
}

func (db Dao) GetCategory(id int) (Category, error) {
	var c Category
	err := db.Db.QueryRow("SELECT id, name, list_id, created_at, updated_at FROM categories WHERE id = ?;", id).Scan(&c.Id, &c.Name, &c.ListId, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		log.Print(err)
		return Category{}, err
	}

	c.Items, err = db.getItemsForCategory(c)
	if err != nil {
		log.Print(err)
	}
	return c, err
}

func (db Dao) CreateCategory(body []byte) (Category, error) {
	var category Category

	unmarshalErr := json.Unmarshal(body, &category)
	if unmarshalErr != nil {
		log.Print(unmarshalErr)
		return Category{}, unmarshalErr
	}

	if category.Name == "" {
		err := ValidationError{Message: "Category name must not be blank"}
		return Category{}, err
	}

	stmt, prepErr := db.Db.Prepare("INSERT INTO categories(name, list_id, created_at) VALUES(?,?, NOW())")
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
	queryErr := db.Db.QueryRow("SELECT id, name, list_id, created_at, updated_at FROM categories WHERE id=?;", lastId).Scan(&category.Id, &category.Name, &category.ListId, &category.CreatedAt, &category.UpdatedAt)
	if queryErr != nil {
		log.Print(queryErr)
		return Category{}, queryErr
	}
	return category, nil
}

func (db Dao) UpdateCategory(id int, body []byte) (Category, error) {
	var category Category
	unmarshalErr := json.Unmarshal(body, &category)
	log.Println("Unmarshal")
	log.Println(body)
	if unmarshalErr != nil {
		log.Print(unmarshalErr)
		return Category{}, unmarshalErr
	}
	log.Println("Body: ", body)
	log.Println("Category: ", category)

	var nameValueBuffer bytes.Buffer
	// TODO: Don't get pwned buy SQL injection
	if category.Name != "" {
		s := fmt.Sprintf("%s = '%s' ", "name", category.Name)
		nameValueBuffer.WriteString(s)
	}

	sql := fmt.Sprintf("UPDATE categories SET %sWHERE id = %d;", nameValueBuffer.String(), id)
	log.Println(sql)
	res, updateErr := db.Db.Exec(sql)
	if updateErr != nil {
		log.Print(updateErr)
		return Category{}, updateErr
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Print(err)
		return Category{}, err
	}
	if rowsAffected == 0 {
		return Category{}, fmt.Errorf("Unable to update category with ID %d", id)
	}

	queryErr := db.Db.QueryRow("SELECT id, name, list_id, created_at, updated_at FROM categories WHERE id = ?;", id).Scan(&category.Id, &category.Name, &category.ListId, &category.CreatedAt, &category.UpdatedAt)
	if queryErr != nil {
		log.Print(queryErr)
		return Category{}, queryErr
	}
	return category, nil
}

// TODO: Cascading delete for items?
func (db Dao) DeleteCategory(id int) (bool, error) {
	if id <= 0 {
		return false, fmt.Errorf("Bad category ID: %d", id)
	}

	stmt, err := db.Db.Prepare("DELETE FROM categories WHERE id = ?;")
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

func (dao Dao) getCategoriesForList(l List) ([]Category, error) {
	rows, err := dao.Db.Query("select id, name, list_id, created_at, updated_at from categories where list_id = ?;", l.Id)
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

		c.Items, err = dao.GetItemsForCategoryAndList(l, c)
		if err != nil {
			log.Print(err)
		}

		foundCategories = append(foundCategories, c)
	}
	return foundCategories, err
}

func printStringSlice(s []string) {
	fmt.Print("printStringSlice")
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
