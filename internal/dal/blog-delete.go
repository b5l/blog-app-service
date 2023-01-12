package dal

import (
	"blog-app-service/internal/errorx"
	"blog-app-service/internal/pkg/dto"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type BlogDeleteDAO interface {
	GetBlogDelete(ctx context.Context, id string) (dto.BlogDeleteResponseBody, *errorx.Error)
}

type blogDeleteDAO struct {
	db *sqlx.DB
}

func (u *blogDeleteDAO) GetBlogDelete(ctx context.Context, id string) (dto.BlogDeleteResponseBody, *errorx.Error) {
	var deleted string
	if err := u.db.Get(&deleted,
		`Delete FROM blog_posts WHERE id = $1`, id,
	); err != nil {
		if err == sql.ErrNoRows {
			return dto.BlogDeleteResponseBody{IsDeleted: true}, nil
		} else {
			return dto.BlogDeleteResponseBody{}, &errorx.Error{
				Message:    err.Error(),
				Details:    "",
				StatusCode: 500,
			}
		}
	}

	return dto.BlogDeleteResponseBody{}, nil
}

// type check
var _ BlogDeleteDAO = &blogDeleteDAO{}

func NewBlogDeleteDAO(db *sqlx.DB) BlogDeleteDAO {
	return &blogDeleteDAO{
		db: db,
	}
}
