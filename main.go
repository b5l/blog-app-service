package main

import (
	blogCreate "blog-app-service/internal/api/blog_create"
	blogDelete "blog-app-service/internal/api/blog_delete"
	blogDetails "blog-app-service/internal/api/blog_details"
	blogEdit "blog-app-service/internal/api/blog_edit"
	blogPosts "blog-app-service/internal/api/blog_posts"
	"blog-app-service/internal/api/login"
	signUp "blog-app-service/internal/api/sign_up"
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
	signUp := &signUp.Handler{
		SignUpDAO: dal.NewSignUpDAO(DB),
	}
	blogPosts := &blogPosts.Handler{
		BlogPostsDAO: dal.NewBlogPostsDAO(DB),
	}
	blogDetails := &blogDetails.Handler{
		BlogDetailsDAO: dal.NewBlogDetailsDAO(DB),
	}
	blogCreate := &blogCreate.Handler{
		BlogCreateEditDAO: dal.NewBlogCreateEditDAO(DB),
	}
	blogEdit := &blogEdit.Handler{
		BlogCreateEditDAO: dal.NewBlogCreateEditDAO(DB),
	}
	blogDelete := &blogDelete.Handler{
		BlogDeleteDAO: dal.NewBlogDeleteDAO(DB),
	}

	router := gin.Default()
	api := router.Group("/api")
	api.POST("/login", login.LoginHandler)
	api.POST("/signUp", signUp.SignUpHandler)
	api.GET("/blogPosts", blogPosts.BlogPostsHandler)
	api.GET("/blogPost/:id", blogDetails.BlogDetailsHandler)
	api.POST("/blogCreate", blogCreate.BlogCreateHandler)
	api.PUT("/blogPost/:id", blogEdit.BlogEditHandler)
	api.DELETE("/blogPost/:id", blogDelete.BlogDeleteHandler)

	router.Run("localhost:8080")
}
