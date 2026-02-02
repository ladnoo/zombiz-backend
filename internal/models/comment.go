package models

import "time"

type Comment struct {
	ID          int       `json:"id"`
	PostID      int       `json:"post_id"`
	UserID      int       `json:"user_id"`
	Nickname    string    `json:"nickname"`
	EmojiAvatar string    `json:"emoji_avatar"`
	Text        string    `json:"text"`
	ImageURLs   []string  `json:"image_urls"`
	CreatedAt   time.Time `json:"created_at"`
}
