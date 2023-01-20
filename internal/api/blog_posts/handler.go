package blogPosts

import (
	"blog-app-service/internal/errorx"
	"blog-app-service/internal/pkg/dto"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DAO interface {
	GetBlogPosts(ctx context.Context) ([]dto.BlogPostsObject, *errorx.Error)
}

type Handler struct {
	BlogPostsDAO DAO
}

func (h *Handler) BlogPostsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogPosts, errx := h.BlogPostsDAO.GetBlogPosts(c)

		if errx != nil {
			c.JSON(errx.StatusCode, errx)
		} else {
			c.JSON(http.StatusOK, &dto.BlogPostsResponseBody{Data: blogPosts})
		}
	}
}
