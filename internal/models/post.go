package models

import "time"

type Post struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Nickname    string    `json:"nickname"`
	EmojiAvatar string    `json:"emoji_avatar"`
	Text        string    `json:"text"`
	ImageURLs   []string  `json:"image_urls"`
	CreatedAt   time.Time `json:"created_at"`
}
