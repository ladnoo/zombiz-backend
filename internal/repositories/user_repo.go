package repositories

import (
	"zombiz/internal/config"
	"zombiz/internal/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(nickname, emojiAvatar string) (*models.User, error) {
	query := `
	INSERT INTO users (nickname, emoji_avatar)
	VALUES ($1, $2)
	RETURNING id, nickname, emoji_avatar, created_at;
	`

	var user models.User

	err := config.DB.QueryRow(query, nickname, emojiAvatar).Scan(
		&user.ID,
		&user.Nickname,
		&user.EmojiAvatar,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByNickname(nickname string) (*models.User, error) {
	query := `
		SELECT id, nickname, emoji_avatar, created_at
		FROM users
		WHERE nickname = $1
	`

	var user models.User

	err := config.DB.QueryRow(query, nickname).Scan(
		&user.ID,
		&user.Nickname,
		&user.EmojiAvatar,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByID(userID int) (*models.User, error) {
	query := `
		SELECT id, nickname, emoji_avatar, created_at
		FROM users
		WHERE id = $1
	`

	var user models.User

	err := config.DB.QueryRow(query, userID).Scan(
		&user.ID,
		&user.Nickname,
		&user.EmojiAvatar,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	query := `
		SELECT id, nickname, emoji_avatar, created_at
		FROM users
		ORDER BY created_at DESC
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.Nickname,
			&user.EmojiAvatar,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
