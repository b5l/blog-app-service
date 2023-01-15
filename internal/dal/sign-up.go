package dal

import (
	"blog-app-service/internal/errorx"
	model "blog-app-service/internal/model/db"
	"blog-app-service/internal/pkg/dto"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type SignUpDAO interface {
	PostUser(ctx context.Context, username string, password string) (dto.SignUpResponseBody, *errorx.Error)
}

type signUpDAO struct {
	db *sqlx.DB
}

func (u *signUpDAO) PostUser(ctx context.Context, username string, password string) (dto.SignUpResponseBody, *errorx.Error) {
	signUp := model.LoginSignUp{
		Username: username,
		Password: password,
	}

	var getUser string
	err := u.db.Get(&getUser, "SELECT username FROM users WHERE username=$1 LIMIT 1", username)

	switch err {
	case nil:
		return dto.SignUpResponseBody{UserTaken: true, IsSuccessful: false}, nil

	case sql.ErrNoRows:
		_, dbErr := u.db.NamedExec("INSERT INTO users (username,password) VALUES (:username, :password)", signUp)
		if dbErr == nil {
			return dto.SignUpResponseBody{UserTaken: false, IsSuccessful: true}, nil
		} else {
			return dto.SignUpResponseBody{}, &errorx.Error{
				Message:    err.Error(),
				Details:    "",
				StatusCode: 500,
			}
		}
	}

	return dto.SignUpResponseBody{}, nil
}

// type check
var _ SignUpDAO = &signUpDAO{}

func NewSignUpDAO(db *sqlx.DB) SignUpDAO {
	return &signUpDAO{
		db: db,
	}
}
