package domain

type BroadCaster struct {
	ID         int    `json:"id"`
	TwitchID   string `json:"twitch_id"`
	UserID     int    `json:"user_id"`
	Name       string `json:"name"`
	ImageURL   string `json:"image_url"`
	ClipCursor string `json:"clip_cursor"`
}
