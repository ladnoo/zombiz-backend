package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"zombiz/internal/config"
	"zombiz/internal/handlers"
)

// Точка входа проекта

func main() {
	config.InitDB()
	defer config.DB.Close()

	r := gin.Default()

	// Главная страница
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Привет!",
			"status":  "ok",
		})
	})

	users := r.Group("/users")
	{
		users.GET("", handlers.GetUsers)
		users.POST("", handlers.CreateUser)
	}

	//Группа эндпоинтов для постов
	posts := r.Group("/posts")
	{
		posts.GET("", handlers.GetPosts)    // Получить все посты
		posts.POST("", handlers.CreatePost) // Создать посты
		posts.GET("/:id", handlers.GetPostByID)
		posts.GET("/user/:user_id", handlers.GetPostByUserID)
	}

	comments := r.Group("/comments")
	{
		comments.GET("/post/:post_id", handlers.GetCommentsByPost)
		comments.POST("/post/:post_id", handlers.CreateComment)
		comments.GET("/:id", handlers.GetCommentByID)
	}

	log.Println("Сревер запущен на порту 8080")
	r.Run(":8080")
}
