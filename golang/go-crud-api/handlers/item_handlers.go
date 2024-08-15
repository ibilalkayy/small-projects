package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ibilalkayy/go-crud-api/db"
	"github.com/ibilalkayy/go-crud-api/models"
)

// CreateItem handles creating a new item
func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	item.Name = r.FormValue("name")
	item.Description = r.FormValue("description")

	_, err := db.DB.Exec(context.Background(), "INSERT INTO items (name, description) VALUES ($1, $2)", item.Name, item.Description)
	if err != nil {
		log.Printf("Error creating item: %v\n", err)
		http.Error(w, "Unable to create item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Item created")
}

// GetItems handles retrieving all items
func GetItems(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(context.Background(), "SELECT id, name, description FROM items")
	if err != nil {
		http.Error(w, "Unable to retrieve items", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		rows.Scan(&item.ID, &item.Name, &item.Description)
		items = append(items, item)
	}

	json.NewEncoder(w).Encode(items)
}

// UpdateItem handles updating an existing item
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	name := r.FormValue("name")
	description := r.FormValue("description")

	_, err := db.DB.Exec(context.Background(), "UPDATE items SET name=$1, description=$2 WHERE id=$3", name, description, id)
	if err != nil {
		http.Error(w, "Unable to update item", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Item updated")
}

// DeleteItem handles deleting an item
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))

	_, err := db.DB.Exec(context.Background(), "DELETE FROM items WHERE id=$1", id)
	if err != nil {
		http.Error(w, "Unable to delete item", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Item deleted")
}
