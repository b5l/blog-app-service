package blogPosts

import (
	"blog-app-service/internal/errorx"
	"blog-app-service/internal/pkg/dto"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DAO interface {
	GetBlogDetails(ctx context.Context, id string) (dto.BlogDetailsResponseBody, *errorx.Error)
}

type Handler struct {
	BlogDetailsDAO DAO
}

func (h *Handler) BlogDetailsHandler(c *gin.Context) {
	id := c.Param("id")

	blogDetails, errx := h.BlogDetailsDAO.GetBlogDetails(c, id)

	if errx != nil {
		c.JSON(errx.StatusCode, errx)
	} else {
		c.JSON(http.StatusOK, blogDetails)
	}
}
