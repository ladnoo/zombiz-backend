package repositories

import (
	"github.com/lib/pq"
	"zombiz/internal/config"
	"zombiz/internal/models"
)

type PostRepository struct{}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func (r *PostRepository) Create(userID int, text string, imageUrls []string) (*models.Post, error) {
	query := `
		INSERT INTO posts (user_id, text, image_urls)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, text, image_urls, created_at;
	`

	var post models.Post
	var imageURLsPQ pq.StringArray

	err := config.DB.QueryRow(query, userID, text, imageUrls).Scan(
		&post.ID,
		&post.UserID,
		&post.Text,
		&imageURLsPQ,
		&post.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	post.ImageURLs = imageURLsPQ
	return &post, nil
}

func (r *PostRepository) GetAll() ([]models.Post, error) {
	query := `
		SELECT p.id, p.user_id, p.text, p.image_urls, p.created_at,
			u.nickname, u.emoji_avatar
		FROM posts p
		JOIN users u ON p.user_id = u.id
		ORDER BY p.created_at DESC
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		var imageURLsPQ pq.StringArray
		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Text,
			&imageURLsPQ,
			&post.CreatedAt,
			&post.Nickname,
			&post.EmojiAvatar,
		)
		if err != nil {
			return nil, err
		}
		post.ImageURLs = imageURLsPQ
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostRepository) GetByID(postID int) (*models.Post, error) {
	query := `SELECT p.id, p.user_id, p.text, p.image_urls, p.created_at,
		u.nickname, u.emoji_avatar
		FROM posts p
		JOIN users u ON p.user_id = u.id
		WHERE p.id = $1;
	`
	var post models.Post
	var imageURLsPQ pq.StringArray
	err := config.DB.QueryRow(query, postID).Scan(
		&post.ID,
		&post.UserID,
		&post.Text,
		&imageURLsPQ,
		&post.CreatedAt,
		&post.Nickname,
		&post.EmojiAvatar,
	)
	if err != nil {
		return nil, err
	}
	post.ImageURLs = imageURLsPQ
	return &post, nil
}
