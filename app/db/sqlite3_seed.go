package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var seeds = `
INSERT INTO users(twitch_id, name, email, image_url, token) 
VALUES('123456789', 'naosuke884', 'naosuke884@example.jp', 'https://example.jp/image.png', 'token');

INSERT INTO broadcasters(twitch_id, user_id, name, image_url, clip_cursor)
VALUES('123456789', 1, 'naosuke884', 'https://example.jp/image.png', 'cursor');

INSERT INTO clips(twitch_id, broadcaster_id, title, thumbnail_url, embed_url, clip_url, view_count)
VALUES('123456789', 1, 'title', 'https://example.jp/image.png', 'https://example.jp/embed', 'https://example.jp/clip', 100);
`

func createSqliteSeed(db *sql.DB) {
	_, err := db.Exec(seeds)
	if err != nil {
		log.Fatal(err)
	}
}
