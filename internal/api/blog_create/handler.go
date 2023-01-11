package blogCreate

import (
	"blog-app-service/internal/errorx"
	model "blog-app-service/internal/model/db"
	"blog-app-service/internal/pkg/dto"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DAO interface {
	BlogCreateEdit(ctx context.Context, id int, title string, shortDescription string, longDescription string) (*dto.BlogCreateEditResponseBody, *errorx.Error)
}

type Handler struct {
	BlogCreateEditDAO DAO
}

func (h *Handler) BlogCreateHandler(c *gin.Context) {
	var getDetails model.BlogDetails

	if err := c.BindJSON(&getDetails); err != nil {
		return
	}

	isSuccess, errx := h.BlogCreateEditDAO.BlogCreateEdit(c, 0, getDetails.Title, getDetails.Type, getDetails.Description)

	if errx != nil {
		c.JSON(errx.StatusCode, errx)
	} else {
		c.JSON(http.StatusOK, isSuccess)
	}
}
