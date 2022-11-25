package domain

type Clip struct {
	ID            int    `json:"id"`
	TwitchID      string `json:"twitch_id"`
	BroadCasterID int    `json:"broadcaster_id"`
	Title         string `json:"title"`
	ThumbnailURL  string `json:"thumbnail_url"`
	EmbedURL      string `json:"embed_url"`
	ClipURL       string `json:"clip_url"`
	ViewCount     int    `json:"view_count"`
}
