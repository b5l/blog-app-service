package main

import (
	blogPosts "blog-app-service/internal/api/blog_posts"
	"blog-app-service/internal/api/login"
	"blog-app-service/internal/dal"
	"blog-app-service/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	DB, err := database.Init()
	if err != nil {
		panic(err)
	}

	login := &login.Handler{
		LoginDAO: dal.NewLoginDAO(DB),
	}

	router := gin.Default()
	api := router.Group("/api")
	api.POST("/login", login.LoginHandler)
	api.GET("/blogpost", blogPosts.BlogPostsHandler)

	router.Run("localhost:8080")
}
