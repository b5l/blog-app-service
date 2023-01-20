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
	GetBlogCreateEdit(ctx context.Context, Id int, Title string, Type string, Description string) (*dto.BlogCreateEditResponseBody, *errorx.Error)
}

type blogCreateEditDAO struct {
	db *sqlx.DB
}

func (u *blogCreateEditDAO) GetBlogCreateEdit(ctx context.Context, Id int, Title string, Type string, Description string) (*dto.BlogCreateEditResponseBody, *errorx.Error) {
	var results *dto.BlogCreateEditResponseBody

	if Title == "" {
		if err := u.db.Get(&Title,
			`SELECT title FROM blog_posts WHERE id=$1 LIMIT 1`, Id,
		); err != nil {
			Title = ""
		}
	}

	if Type == "" {
		if err := u.db.Get(&Type,
			`SELECT type FROM blog_posts WHERE id=$1 LIMIT 1`, Id,
		); err != nil {
			Type = ""
		}
	}

	if Description == "" {
		if err := u.db.Get(&Description,
			`SELECT description FROM blog_posts WHERE id=$1 LIMIT 1`, Id,
		); err != nil {
			Description = ""
		}
	}

	create := model.BlogCreate{
		Title:       Title,
		Type:        Type,
		Description: Description,
	}
	edit := model.BlogDetails{
		Id:          Id,
		Title:       Title,
		Type:        Type,
		Description: Description,
	}

	var blogID int
	err := u.db.Get(&blogID, "SELECT id FROM blog_posts WHERE id=$1 LIMIT 1", Id)

	switch err {
	case nil:
		_, dbErr := u.db.NamedExec("UPDATE blog_posts SET title = :title, type = :type, description = :description WHERE id=:id", edit)
		if dbErr != nil {
			fmt.Println(dbErr)

			return nil, &errorx.Error{
				Message:    dbErr.Error(),
				Details:    "",
				StatusCode: 404,
			}
		}
	case sql.ErrNoRows:
		_, dbErr := u.db.NamedExec("INSERT INTO blog_posts (title, type, description) VALUES (:title, :type, :description)", create)
		if dbErr != nil {
			fmt.Println(dbErr)

			return nil, &errorx.Error{
				Message:    dbErr.Error(),
				Details:    "",
				StatusCode: 500,
			}
		}
	}

	results = &dto.BlogCreateEditResponseBody{IsSuccessful: true}

	return results, nil
}

// type check
var _ BlogCreateEditDAO = &blogCreateEditDAO{}

func NewBlogCreateEditDAO(db *sqlx.DB) BlogCreateEditDAO {
	return &blogCreateEditDAO{
		db: db,
	}
}
