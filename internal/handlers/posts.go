package handlers

import "github.com/gin-gonic/gin"

func GetPosts(c *gin.Context) {
	c.JSON(200, gin.H{
		"posts": []gin.H{},
	})
}

func CreatePost(c *gin.Context) {
	var postData struct {
		Text string `json:"text"`
	}

	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(400, gin.H{"error": "Неверный формат данных"})
		return
	}

	c.JSON(201, gin.H{
		"message": "Пост создан",
		"text":    postData.Text,
	})
}
