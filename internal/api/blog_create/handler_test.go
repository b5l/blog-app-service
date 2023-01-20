package blogCreate

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

func TestHandler_BlogCreate(t *testing.T) {
	type fields struct {
		BlogCreateEditDAO DAO
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
				BlogCreateEditDAO: &mockDAO{
					blogCreate: dto.BlogCreateEditResponseBody{
						IsSuccessful: true,
					},
				},
			},
			status: http.StatusOK,
			body: &dto.BlogCreateEditResponseBody{
				IsSuccessful: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				BlogCreateEditDAO: tt.fields.BlogCreateEditDAO,
			}

			gin.SetMode(gin.ReleaseMode)
			r := gin.Default()
			api := r.Group("/api")
			api.POST("/blogCreate", h.BlogCreateHandler())

			jsonValue, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("POST", "/api/blogCreate", bytes.NewBuffer(jsonValue))
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Result().StatusCode != tt.status {
				t.Errorf("Handler.BlogCreate() = %v, want %v", w.Result().StatusCode, tt.status)
				return
			}

			var bodyMap interface{}
			if err := json.NewDecoder(w.Body).Decode(&bodyMap); err != nil {
				t.Errorf("Handler.BlogCreate() = json body decode error %v", err)
				return
			}

			var wantBodyMap interface{}
			if enc, err := json.Marshal(tt.body); err == nil {
				_ = json.Unmarshal(enc, &wantBodyMap)
			}

			if !reflect.DeepEqual(bodyMap, wantBodyMap) {
				t.Errorf("Handler.BlogCreate() = %v, want %v", bodyMap, wantBodyMap)
			}
		})
	}
}
