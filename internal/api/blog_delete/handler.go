package blogCreate

import (
	"blog-app-service/internal/errorx"
	"blog-app-service/internal/pkg/dto"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DAO interface {
	GetBlogDelete(ctx context.Context, id string) (dto.BlogDeleteResponseBody, *errorx.Error)
}

type Handler struct {
	BlogDeleteDAO DAO
}

func (h *Handler) BlogDeleteHandler(c *gin.Context) {
	id := c.Param("id")

	isDeleted, errx := h.BlogDeleteDAO.GetBlogDelete(c, id)

	if errx != nil {
		c.JSON(errx.StatusCode, errx)
	} else {
		c.JSON(http.StatusOK, isDeleted)
	}
}
