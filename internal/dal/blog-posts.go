package dal

import (
	"blog-app-service/internal/errorx"
	"blog-app-service/internal/pkg/dto"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type BlogPostsDAO interface {
	GetBlogPosts(ctx context.Context) ([]dto.BlogPostsObject, *errorx.Error)
}

type blogPostsDAO struct {
	db *sqlx.DB
}

func (u *blogPostsDAO) GetBlogPosts(ctx context.Context) ([]dto.BlogPostsObject, *errorx.Error) {
	results := []dto.BlogPostsObject{}

	if err := u.db.Select(&results,
		`SELECT id, title, short_description FROM blog_posts`,
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

	return results, nil
}

// type check
var _ BlogPostsDAO = &blogPostsDAO{}

func NewBlogPostsDAO(db *sqlx.DB) BlogPostsDAO {
	return &blogPostsDAO{
		db: db,
	}
}
