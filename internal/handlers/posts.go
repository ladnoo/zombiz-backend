package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPosts(c *gin.Context) {
	posts, err := postRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при получении постов",
		})
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func CreatePost(c *gin.Context) {
	var postData struct {
		UserID    int      `json:"user_id" binding:"required"`
		Text      string   `json:"text"`
		ImageURLs []string `json:"image_urls"`
	}

	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	user, err := userRepo.GetByID(postData.UserID)
	if err != nil || user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Пользователь не найден"})
		return
	}

	post, err := postRepo.Create(postData.UserID, postData.Text, postData.ImageURLs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании поста: " + err.Error()})
		return
	}

	post.Nickname = user.Nickname
	post.EmojiAvatar = user.EmojiAvatar

	c.JSON(http.StatusCreated, gin.H{
		"message": "Пост создан",
		"post":    post,
	})
}

func GetPostByID(c *gin.Context) {
	postID := c.Param("id")

	var id int
	_, err := fmt.Sscan(postID, &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	post, err := postRepo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Пост не найден"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}
