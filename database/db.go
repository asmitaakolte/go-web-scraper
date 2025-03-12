package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

// DB instance
var DB *sql.DB

// InitDB initializes the SQLite database
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "scraper.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	// Create table
	query := `
	CREATE TABLE IF NOT EXISTS articles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		link TEXT
	);`
	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	log.Println("Database initialized successfully")
}

// SaveArticle inserts data into the database
func SaveArticle(title, link string) {
	_, err := DB.Exec("INSERT INTO articles (title, link) VALUES (?, ?)", title, link)
	if err != nil {
		log.Println("Error inserting data:", err)
	}
}
