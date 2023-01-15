package login

import (
	"blog-app-service/internal/errorx"
	model "blog-app-service/internal/model/db"
	"blog-app-service/internal/pkg/dto"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DAO interface {
	PostUser(ctx context.Context, username string, password string) (dto.SignUpResponseBody, *errorx.Error)
}

type Handler struct {
	SignUpDAO DAO
}

func (h *Handler) SignUpHandler(c *gin.Context) {
	var postUser model.LoginSignUp

	if err := c.BindJSON(&postUser); err != nil {
		return
	}

	isSuccessful, errx := h.SignUpDAO.PostUser(c, postUser.Username, postUser.Password)

	if errx != nil {
		c.JSON(errx.StatusCode, errx)
	} else {
		c.JSON(http.StatusOK, isSuccessful)
	}
}
