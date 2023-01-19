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

func getDBForBlogCreateEdit() *sqlx.DB {
	db, err := database.Init()
	if err != nil {
		panic(err)
	}
	return db
}

func Test_BlogCreateEdit(t *testing.T) {
	type fields struct {
		DB    *sqlx.DB
		setup []string
	}
	type args struct {
		ctx         context.Context
		Id          int
		Title       string
		Type        string
		Description string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.BlogCreateEditResponseBody
		wantErr bool
	}{
		{
			name: "Create",
			fields: fields{
				DB: getDBForBlogCreateEdit(),
			},
			args: args{
				Title:       "Create",
				Type:        "Type",
				Description: "Desc",
			},
			want: &dto.BlogCreateEditResponseBody{IsSuccessful: true},
		},
		{
			name: "Edit",
			fields: fields{
				DB: getDBForBlogCreateEdit(),
				setup: []string{`
					INSERT INTO blog_posts (id, title, type, description)
					VALUES (2,  'Edit', 'Type', 'Desc');`,
				},
			},
			args: args{
				Id:          2,
				Title:       "Edit",
				Type:        "Type",
				Description: "Desc",
			},
			want: &dto.BlogCreateEditResponseBody{IsSuccessful: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, s := range tt.fields.setup {
				tt.fields.DB.MustExec(s)
			}
			u := dal.NewBlogCreateEditDAO(tt.fields.DB)

			got, err := u.BlogCreateEdit(context.Background(), tt.args.Id, tt.args.Title, tt.args.Type, tt.args.Description)
			if (err != nil) != tt.wantErr {
				t.Errorf("BlogCreateEdit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BlogCreateEdit = %v, want %v", got, tt.want)
			}
			tt.fields.DB.Query(`Delete FROM blog_posts WHERE title='Create'`)
			tt.fields.DB.Query(`Delete FROM blog_posts WHERE id=2`)
		})
	}
}
