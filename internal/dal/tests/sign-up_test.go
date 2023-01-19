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

func getDBForSignUp() *sqlx.DB {
	db, err := database.Init()
	if err != nil {
		panic(err)
	}
	return db
}

func Test_SignUp(t *testing.T) {
	type fields struct {
		DB    *sqlx.DB
		setup []string
	}
	type args struct {
		ctx      context.Context
		Username string
		Password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.SignUpResponseBody
		wantErr bool
	}{
		{
			name: "Fetch",
			fields: fields{
				DB: getDBForSignUp(),
			},
			args: args{
				Username: "Test username",
				Password: "Test password",
			},
			want: dto.SignUpResponseBody{IsSuccessful: true},
		},
		{
			name: "User taken",
			fields: fields{
				DB: getDBForSignUp(),
				setup: []string{`
				INSERT INTO users (username, password)
				VALUES ('Test username',  'Test password');`,
				},
			},
			args: args{
				Username: "Test username",
				Password: "Test password",
			},
			want: dto.SignUpResponseBody{UserTaken: true, IsSuccessful: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, s := range tt.fields.setup {
				tt.fields.DB.MustExec(s)
			}
			u := dal.NewSignUpDAO(tt.fields.DB)

			got, err := u.PostUser(context.Background(), tt.args.Username, tt.args.Password)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostUser = %v, want %v", got, tt.want)
			}
			tt.fields.DB.Query(`Delete FROM users WHERE username='Test username' AND password='Test password'`)
		})
	}
}
