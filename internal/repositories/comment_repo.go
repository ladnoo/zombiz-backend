package repositories

import (
	"zombiz/internal/config"
	"zombiz/internal/models"
)

type CommentRepository struct{}

func NewCommentRepository() *CommentRepository {
	return &CommentRepository{}
}

func (r *CommentRepository) Create(postID, userID int, text string, imageURLs []string) (*models.Comment, error) {
	query := `
		INSERT INTO comments(post_id, user_id, text, image_urls)
		VALUES ($1, $2, $3, $4)
		RETURNING id, post_id, user_id, text, image_urls, created_at
	`

	var comment models.Comment

	err := config.DB.QueryRow(query, postID, userID, text, imageURLs).Scan(
		&comment.ID,
		&comment.PostID,
		&comment.UserID,
		&comment.Text,
		&comment.ImageURLs,
		&comment.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &comment, nil

}

func (r *CommentRepository) GetByPostID(postID int) ([]models.Comment, error) {
	query := `SELECT c.id, c.post_id, c.user_id, c.text, c.image_urls, c.created_at
			u.nickname, u.emoji_avatar
		FROM comments c
		JOIN users u ON c.user_id = u.id
		WHERE c.post_id = $1
	`

	rows, err := config.DB.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.UserID,
			&comment.Text,
			&comment.ImageURLs,
			&comment.CreatedAt,
			&comment.Nickname,
			&comment.EmojiAvatar,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (r *CommentRepository) GetAll() ([]models.Comment, error) {
	query := `SELECT c.id, c.post_id, c.user_id, c.text, c.image_urls, c.created_at,
			u.nickname, u.emoji_avatar
		FROM comments c
		JOIN users u ON c.user_id = u.id
		ORDER BY c.created_at DESC
	`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.UserID,
			&comment.Text,
			&comment.ImageURLs,
			&comment.CreatedAt,
			&comment.Nickname,
			&comment.EmojiAvatar,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
