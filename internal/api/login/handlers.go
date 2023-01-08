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
	GetUser(ctx context.Context, username string, password string) (*dto.LoginResponseBody, *errorx.Error)
}

type Handler struct {
	LoginDAO DAO
}

func (h *Handler) LoginHandler(c *gin.Context) {
	var getUser model.Login

	if err := c.BindJSON(&getUser); err != nil {
		return
	}

	isAuth, errx := h.LoginDAO.GetUser(c, getUser.Username, getUser.Password)

	if errx != nil {
		c.JSON(errx.StatusCode, errx)
	} else {
		c.JSON(http.StatusOK, isAuth)
	}
}
