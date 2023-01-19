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

func getDBForLogin() *sqlx.DB {
	db, err := database.Init()
	if err != nil {
		panic(err)
	}
	return db
}

func Test_Login(t *testing.T) {
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
		want    *dto.LoginResponseBody
		wantErr bool
	}{
		{
			name: "Fetch",
			fields: fields{
				DB: getDBForLogin(),
				setup: []string{`
				INSERT INTO users (username, password)
				VALUES ('Test username',  'Test password');`,
				},
			},
			args: args{
				Username: "Test username",
				Password: "Test password",
			},
			want: &dto.LoginResponseBody{IsAuth: true},
		},
		{
			name: "SQL No rows error",
			fields: fields{
				DB: getDBForLogin(),
			},
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, s := range tt.fields.setup {
				tt.fields.DB.MustExec(s)
			}
			u := dal.NewLoginDAO(tt.fields.DB)

			got, err := u.GetUser(context.Background(), tt.args.Username, tt.args.Password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser = %v, want %v", got, tt.want)
			}
			tt.fields.DB.Query(`Delete FROM users WHERE username='Test username' AND password='Test password'`)
		})
	}
}
