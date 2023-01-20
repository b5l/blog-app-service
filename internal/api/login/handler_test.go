package login

import (
	"blog-app-service/internal/pkg/dto"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandler_Login(t *testing.T) {
	type fields struct {
		LoginDAO DAO
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		status int
		body   interface{}
	}{
		{
			name: "success",
			fields: fields{
				LoginDAO: &mockDAO{
					login: dto.LoginResponseBody{
						IsAuth: true,
					},
				},
			},
			status: http.StatusOK,
			body: &dto.LoginResponseBody{
				IsAuth: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				LoginDAO: tt.fields.LoginDAO,
			}

			gin.SetMode(gin.ReleaseMode)
			r := gin.Default()
			api := r.Group("/api")
			api.POST("/login", h.LoginHandler())

			jsonValue, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonValue))
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Result().StatusCode != tt.status {
				t.Errorf("Handler.Login() = %v, want %v", w.Result().StatusCode, tt.status)
				return
			}

			var bodyMap interface{}
			if err := json.NewDecoder(w.Body).Decode(&bodyMap); err != nil {
				t.Errorf("Handler.Login() = json body decode error %v", err)
				return
			}

			var wantBodyMap interface{}
			if enc, err := json.Marshal(tt.body); err == nil {
				_ = json.Unmarshal(enc, &wantBodyMap)
			}

			if !reflect.DeepEqual(bodyMap, wantBodyMap) {
				t.Errorf("Handler.Login() = %v, want %v", bodyMap, wantBodyMap)
			}
		})
	}
}
