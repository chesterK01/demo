package main

import (
	"database/sql"
	"demo/db"
	"fmt"
	"log"
	"net/http"
	"os"
)

var database *sql.DB

func main() {
	// Khởi tạo database
	database = db.InitDB()
	defer database.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Lấy message từ database
		var message string
		err := database.QueryRow("SELECT message FROM demo_db LIMIT 1").Scan(&message)
		if err != nil {
			log.Printf("Error querying database: %v", err)
			message = "Hello World!"
		}
		fmt.Fprintf(w, message)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on :%s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
