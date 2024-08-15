package main

import (
	"log"
	"net/http"

	"github.com/ibilalkayy/go-crud-api/db"
	"github.com/ibilalkayy/go-crud-api/handlers"
)

func main() {
	db.ConnectDB()

	http.HandleFunc("/create", handlers.CreateItem)
	http.HandleFunc("/items", handlers.GetItems)
	http.HandleFunc("/update", handlers.UpdateItem)
	http.HandleFunc("/delete", handlers.DeleteItem)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
