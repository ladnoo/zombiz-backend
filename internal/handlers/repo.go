package handlers

import "zombiz/internal/repositories"

var (
	userRepo    = repositories.NewUserRepository()
	commentRepo = repositories.NewCommentRepository()
	postRepo    = repositories.NewPostRepository()
)
