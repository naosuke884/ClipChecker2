package domain

type User struct {
	ID       ID     `json:"id"`
	TwitchID string `json:"twitch_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	ImageURL string `json:"image_url"`
	Token    string `json:"token"`
}

type UserRepository interface {
	Store(*User) error
	Find(*ID) (*User, error)
	Update(*User) error
	Delete(*ID) error
}
