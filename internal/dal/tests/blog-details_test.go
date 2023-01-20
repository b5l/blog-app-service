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

func getDBForBlogDetails() *sqlx.DB {
	db, err := database.Init()
	if err != nil {
		panic(err)
	}
	return db
}

func Test_BlogDetails(t *testing.T) {
	type fields struct {
		DB    *sqlx.DB
		setup []string
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.BlogDetailsResponseBody
		wantErr bool
	}{
		{
			name: "Fetch",
			fields: fields{
				DB: getDBForBlogDetails(),
				setup: []string{`
				INSERT INTO blog_posts (id, title, type, description)
				VALUES (1,  'Test title 1', 'Test type 1', 'Test desc 1');
				INSERT INTO blog_posts (id, title, type, description)
				VALUES (2,  'Test title 2', 'Test type 2', 'Test desc 2');
				INSERT INTO blog_posts (id, title, type, description)
				VALUES (3,  'Test title 3', 'Test type 3','Test desc 3');
				`},
			},
			args: args{
				id: "3",
			},
			want: dto.BlogDetailsResponseBody{
				Id:          3,
				Title:       "Test title 3",
				Type:        "Test type 3",
				Description: "Test desc 3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, s := range tt.fields.setup {
				tt.fields.DB.MustExec(s)
			}
			u := dal.NewBlogDetailsDAO(tt.fields.DB)

			got, err := u.GetBlogDetails(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlogDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBlogDetails = %v, want %v", got, tt.want)
			}
			tt.fields.DB.Query(`Delete FROM blog_posts WHERE id=1`)
			tt.fields.DB.Query(`Delete FROM blog_posts WHERE id=2`)
			tt.fields.DB.Query(`Delete FROM blog_posts WHERE id=3`)
		})
	}
}
