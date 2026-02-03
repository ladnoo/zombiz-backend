package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCommentsByPost(c *gin.Context) {
	postID := c.Param("post_id")

	var id int
	_, err := fmt.Sscan(postID, &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID поста"})
		return
	}

	post, err := postRepo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при проверке поста"})
		return
	}
	if post == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Пост не найден",
		})
		return
	}

	comments, err := commentRepo.GetByPostID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при получении комментариев",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count":   len(comments),
		"comment": comments,
	})
}

func CreateComment(c *gin.Context) {
	id := c.Param("post_id")

	var postID int
	_, err := fmt.Sscan(id, &postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Неверный формат ID поста",
		})
		return
	}

	post, err := postRepo.GetByID(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при проверке поста",
		})
	}
	if post == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	var input struct {
		UserID    int      `json:"user_id" binding:"required"`
		Text      string   `json:"text"`
		ImageURLs []string `json:"image_urls"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	user, err := userRepo.GetByID(post.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при проверке пользователя"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	if input.Text == "" && input.ImageURLs == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Коментарий должен иметь текст или изображение"})
		return
	}

	comment, err := commentRepo.Create(postID, input.UserID, input.Text, input.ImageURLs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании комментария"})
		return
	}

	comment.Nickname = user.Nickname
	comment.EmojiAvatar = user.EmojiAvatar
	comment.PostID = postID

	c.JSON(http.StatusCreated, gin.H{
		"message": "Комментарий создан",
		"comment": comment,
	})

}

func GetAllComments(c *gin.Context) {
	comments, err := commentRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении комментариев"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total":    len(comments),
		"comments": comments,
	})
}
