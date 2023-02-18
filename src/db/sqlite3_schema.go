package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	twitch_id TEXT NOT NULL,
	name TEXT NOT NULL,
	email TEXT NOT NULL,
	image_url TEXT NOT NULL,
	token TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS clips (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	twitch_id TEXT NOT NULL,
	broadcaster_id INTEGER NOT NULL,
	title TEXT NOT NULL,
	thumbnail_url TEXT NOT NULL,
	embed_url TEXT NOT NULL,
	clip_url TEXT NOT NULL,
	view_count INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS broadcasters (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	twitch_id TEXT NOT NULL,
	user_id INTEGER NOT NULL,
	name TEXT NOT NULL,
	image_url TEXT NOT NULL,
	clip_cursor TEXT NOT NULL
);
`

func createSqliteSchema(db *sql.DB) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}
