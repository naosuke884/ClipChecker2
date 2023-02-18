package infrastructure

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/naosuke884/ClipChecker2/later/domain"
)

type UserSqlite3Repository struct {
	db *sql.DB
}

func NewUserSqlite3Repository(db *sql.DB) domain.IUserRepository {
	return &UserSqlite3Repository{db: db}
}

func (r *UserSqlite3Repository) Store(user *domain.User) error {
	cmd := `INSERT INTO users(twitch_id, name, email, image_url, token) VALUES(?, ?, ?, ?, ?)`
	params := []any{
		user.TwitchID,
		user.Name,
		user.Email,
		user.ImageURL,
		user.Token,
	}
	_, err := r.db.Exec(cmd, params...)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserSqlite3Repository) Find(id domain.ID) (*domain.User, error) {
	cmd := `SELECT * FROM users WHERE id = ?`
	params := []any{id}
	row := r.db.QueryRow(cmd, params...)
	var user domain.User
	if err := row.Scan(&user.ID, &user.TwitchID, &user.Name, &user.Email, &user.ImageURL, &user.Token); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserSqlite3Repository) Update(user *domain.User) error {
	cmd := `UPDATE users SET twitch_id = ?, name = ?, email = ?, image_url = ?, token = ? WHERE id = ?`
	params := []any{
		user.TwitchID,
		user.Name,
		user.Email,
		user.ImageURL,
		user.Token,
		user.ID,
	}
	_, err := r.db.Exec(cmd, params...)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserSqlite3Repository) Delete(id domain.ID) error {
	cmd := `DELETE FROM users WHERE id = ?`
	params := []any{id}
	_, err := r.db.Exec(cmd, params...)
	if err != nil {
		return err
	}
	return nil
}
