package db

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	// Enable SQLite foreign keys via DSN
	DB, err = sql.Open("sqlite3", "file:api.db?_foreign_keys=on")
	if err != nil {
		panic("Could not connect to database: " + err.Error())
	}

	if err := DB.Ping(); err != nil {
		panic("Could not ping database: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxIdleTime(5 * time.Minute)

	createTable()
}

func createTable() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`
	if _, err := DB.Exec(createUserTable); err != nil {
		panic("Could not create users table: " + err.Error())
	}

	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`
	if _, err := DB.Exec(createEventTable); err != nil {
		panic("Could not create events table: " + err.Error())
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	created_at DATETIME NOT NULL,

	-- prevent duplicate registration for same user+event
	UNIQUE(event_id, user_id),

	FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`
	if _, err := DB.Exec(createRegistrationTable); err != nil {
		panic("Could not create registrations table: " + err.Error())
	}

}
