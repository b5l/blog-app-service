package dal

import (
	"blog-app-service/internal/dal"
	"blog-app-service/internal/database"
	"blog-app-service/internal/pkg/dto"
	"context"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
)

func getDBForBlogPosts() *sqlx.DB {
	db, err := database.Init()
	if err != nil {
		panic(err)
	}
	return db
}

func Test_BlogPosts(t *testing.T) {
	type fields struct {
		DB    *sqlx.DB
		setup []string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		want    []dto.BlogPostsObject
		wantErr bool
	}{
		{
			name: "Fetch",
			fields: fields{
				DB: getDBForBlogPosts(),
				setup: []string{`
				INSERT INTO blog_posts (id, title, type)
				VALUES (1,  'Test title 1', 'Test type 1');
				INSERT INTO blog_posts (id, title, type)
				VALUES (2,  'Test title 2', 'Test type 2');
				INSERT INTO blog_posts (id, title, type)
				VALUES (3,  'Test title 3', 'Test type 3');
				`},
			},
			want: []dto.BlogPostsObject{
				{
					Id:    3,
					Title: "Test title 3",
					Type:  "Test type 3",
				},
				{
					Id:    2,
					Title: "Test title 2",
					Type:  "Test type 2",
				},
				{
					Id:    1,
					Title: "Test title 1",
					Type:  "Test type 1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, s := range tt.fields.setup {
				tt.fields.DB.MustExec(s)
			}
			u := dal.NewBlogPostsDAO(tt.fields.DB)

			got, err := u.GetBlogPosts(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlogPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBlogPosts = %v, want %v", got, tt.want)
			}
			tt.fields.DB.Query(`Delete FROM blog_posts WHERE id=1`)
			tt.fields.DB.Query(`Delete FROM blog_posts WHERE id=2`)
			tt.fields.DB.Query(`Delete FROM blog_posts WHERE id=3`)
		})
	}
}
