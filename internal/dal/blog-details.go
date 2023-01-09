package dal

import (
	"blog-app-service/internal/errorx"
	"blog-app-service/internal/pkg/dto"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type BlogDetailsDAO interface {
	GetBlogDetails(ctx context.Context, id string) (dto.BlogDetailsResponseBody, *errorx.Error)
}

type blogDetailsDAO struct {
	db *sqlx.DB
}

func (u *blogDetailsDAO) GetBlogDetails(ctx context.Context, id string) (dto.BlogDetailsResponseBody, *errorx.Error) {
	results := dto.BlogDetailsResponseBody{}

	if err := u.db.Get(&results,
		`SELECT * FROM blog_posts WHERE id = $1`, id,
	); err != nil {
		if err == sql.ErrNoRows {
			return dto.BlogDetailsResponseBody{}, &errorx.Error{
				Message:    err.Error(),
				Details:    "",
				StatusCode: 404,
			}
		} else {
			return dto.BlogDetailsResponseBody{}, &errorx.Error{
				Message:    err.Error(),
				Details:    "",
				StatusCode: 500,
			}
		}
	}

	return results, nil
}

// type check
var _ BlogDetailsDAO = &blogDetailsDAO{}

func NewBlogDetailsDAO(db *sqlx.DB) BlogDetailsDAO {
	return &blogDetailsDAO{
		db: db,
	}
}
