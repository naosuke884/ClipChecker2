package db

import (
	"database/sql"
	"log"
)

func InitializeSchema(seed bool) {
	db, err := sql.Open("sqlite3", "./db/ClipChecker.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createSqliteSchema(db)
	if seed {
		createSqliteSeed(db)
	}
}
