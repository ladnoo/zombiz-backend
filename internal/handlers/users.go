package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zombiz/internal/repositories"
)

var userRepo = repositories.NewUserRepository()

func CreateUser(c *gin.Context) {
	var input struct {
		Nickname    string `json:"nickname"`
		EmojiAvatar string `json:"emoji_avatar"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	if input.EmojiAvatar == "" {
		input.EmojiAvatar = "😀"
	}

	existingUser, _ := userRepo.GetByNickname(input.Nickname)
	if existingUser != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Пользователь с таким никнеймом уже существует"})
		return
	}

	user, err := userRepo.Create(input.Nickname, input.EmojiAvatar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании пользователя: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Пользователь создан",
		"user":    user,
	})
}

func GetUsers(c *gin.Context) {
	users, err := userRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении пользователей"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
