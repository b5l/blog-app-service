package dal

import (
	"blog-app-service/internal/errorx"
	model "blog-app-service/internal/model/db"
	"blog-app-service/internal/pkg/dto"
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type BlogCreateEditDAO interface {
	BlogCreateEdit(ctx context.Context, id int, title string, shortDescription string, longDescription string) (*dto.BlogCreateEditResponseBody, *errorx.Error)
}

type blogCreateEditDAO struct {
	db *sqlx.DB
}

func (u *blogCreateEditDAO) BlogCreateEdit(ctx context.Context, id int, title string, shortDescription string, longDescription string) (*dto.BlogCreateEditResponseBody, *errorx.Error) {
	var results *dto.BlogCreateEditResponseBody

	if title == "" {
		var getTitle string
		if err := u.db.Get(&getTitle,
			`SELECT title FROM blog_posts WHERE id=$1 LIMIT 1`, id,
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
		title = getTitle
	}

	if shortDescription == "" {
		var getShortDesc string
		if err := u.db.Get(&getShortDesc,
			`SELECT short_description FROM blog_posts WHERE id=$1 LIMIT 1`, id,
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
		shortDescription = getShortDesc
	}

	if longDescription == "" {
		var getLongDesc string
		if err := u.db.Get(&getLongDesc,
			`SELECT long_description FROM blog_posts WHERE id=$1 LIMIT 1`, id,
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
		longDescription = getLongDesc
	}

	create := model.BlogCreate{
		Title:            title,
		ShortDescription: shortDescription,
		LongDescription:  longDescription,
	}
	edit := model.BlogDetails{
		Id:               id,
		Title:            title,
		ShortDescription: shortDescription,
		LongDescription:  longDescription,
	}

	var blogID int
	err := u.db.Get(&blogID, "SELECT id FROM blog_posts WHERE id=$1 LIMIT 1", id)

	switch err {
	case nil:
		_, dbErr := u.db.NamedExec("UPDATE blog_posts SET title = :title, short_description = :short_description, long_description = :long_description WHERE id=:id", edit)
		if dbErr != nil {
			fmt.Println(dbErr)

			return nil, &errorx.Error{
				Message:    dbErr.Error(),
				Details:    "",
				StatusCode: 404,
			}
		}
	case sql.ErrNoRows:
		_, dbErr := u.db.NamedExec("INSERT INTO blog_posts (title, short_description, long_description) VALUES (:title, :short_description, :long_description)", create)
		if dbErr != nil {
			fmt.Println(dbErr)

			return nil, &errorx.Error{
				Message:    dbErr.Error(),
				Details:    "",
				StatusCode: 500,
			}
		}
	}

	results = &dto.BlogCreateEditResponseBody{IsSuccess: true}

	return results, nil
}

// type check
var _ BlogCreateEditDAO = &blogCreateEditDAO{}

func NewBlogCreateEditDAO(db *sqlx.DB) BlogCreateEditDAO {
	return &blogCreateEditDAO{
		db: db,
	}
}
