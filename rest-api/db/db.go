package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Adding underscore to import the package but not using it directly
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

// Create Event table
func createTables() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`
	//#FOREIGN KEY (user_id) REFERENCES users(id)

	_, err := DB.Exec(createEventTable) //Exec return result & error

	if err != nil {
		panic("failed to create event table")
	}
	_, err = DB.Exec(createUserTable)
	if err != nil {
		panic("failed to create user table")
	}
}
