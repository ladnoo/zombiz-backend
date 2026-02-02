package models

import "time"

type User struct {
	ID          int       `json:"id"`
	Nickname    string    `json:"nickname"`
	EmojiAvatar string    `json:"emoji_avatar"`
	CreatedAt   time.Time `json:"created_at"`
}
