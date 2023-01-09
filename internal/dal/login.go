package dal

import (
	"blog-app-service/internal/errorx"
	model "blog-app-service/internal/model/db"
	"blog-app-service/internal/pkg/dto"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type LoginDAO interface {
	GetUser(ctx context.Context, username string, password string) (*dto.LoginResponseBody, *errorx.Error)
}

type loginDAO struct {
	db *sqlx.DB
}

func (u *loginDAO) GetUser(ctx context.Context, username string, password string) (*dto.LoginResponseBody, *errorx.Error) {
	var results *dto.LoginResponseBody
	var user model.Login
	if err := u.db.Get(&user,
		`SELECT * FROM users
		WHERE username = $1
		AND password = $2`,
		username, password,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, &errorx.Error{
				Message:    err.Error(),
				Details:    "",
				StatusCode: 404,
			}
		} else {
			return nil, &errorx.Error{
				Message:    err.Error(),
				Details:    "",
				StatusCode: 500,
			}
		}
	}

	results = &dto.LoginResponseBody{IsAuth: true}

	return results, nil
}

// type check
var _ BlogPostsDAO = &blogPostsDAO{}

func NewLoginDAO(db *sqlx.DB) LoginDAO {
	return &loginDAO{
		db: db,
	}
}
