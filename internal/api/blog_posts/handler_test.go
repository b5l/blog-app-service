package blogPosts

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

func TestHandler_BlogPosts(t *testing.T) {
	type fields struct {
		BlogPostsDAO DAO
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
				BlogPostsDAO: &mockDAO{
					blogPosts: []dto.BlogPostsObject{
						{
							Id:    1,
							Title: "Test Title",
							Type:  "Test Type",
						},
						{
							Id:    2,
							Title: "Test Title 2",
							Type:  "Test Type 2",
						},
						{
							Id:    3,
							Title: "Test Title 3",
							Type:  "Test Type 3",
						},
					},
				},
			},
			status: http.StatusOK,
			body: &dto.BlogPostsResponseBody{
				Data: []dto.BlogPostsObject{

					{
						Id:    1,
						Title: "Test Title",
						Type:  "Test Type",
					},
					{
						Id:    2,
						Title: "Test Title 2",
						Type:  "Test Type 2",
					},
					{
						Id:    3,
						Title: "Test Title 3",
						Type:  "Test Type 3",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				BlogPostsDAO: tt.fields.BlogPostsDAO,
			}

			gin.SetMode(gin.ReleaseMode)
			r := gin.Default()
			api := r.Group("/api")
			api.GET("/blogPosts", h.BlogPostsHandler())

			jsonValue, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("GET", "/api/blogPosts", bytes.NewBuffer(jsonValue))
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Result().StatusCode != tt.status {
				t.Errorf("Handler.BlogPosts() = %v, want %v", w.Result().StatusCode, tt.status)
				return
			}

			var bodyMap interface{}
			if err := json.NewDecoder(w.Body).Decode(&bodyMap); err != nil {
				t.Errorf("Handler.BlogPosts() = json body decode error %v", err)
				return
			}

			var wantBodyMap interface{}
			if enc, err := json.Marshal(tt.body); err == nil {
				_ = json.Unmarshal(enc, &wantBodyMap)
			}

			if !reflect.DeepEqual(bodyMap, wantBodyMap) {
				t.Errorf("Handler.BlogPosts() = %v, want %v", bodyMap, wantBodyMap)
			}
		})
	}
}
