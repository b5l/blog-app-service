package blogDetails

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

func TestHandler_BlogDetails(t *testing.T) {
	type fields struct {
		BlogDetailsDAO DAO
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
				BlogDetailsDAO: &mockDAO{
					blogDetails: dto.BlogDetailsResponseBody{
						Id:          1,
						Title:       "Test Title",
						Type:        "Test Type",
						Description: "Test Desc",
					},
				},
			},
			status: http.StatusOK,
			body: &dto.BlogDetailsResponseBody{
				Id:          1,
				Title:       "Test Title",
				Type:        "Test Type",
				Description: "Test Desc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				BlogDetailsDAO: tt.fields.BlogDetailsDAO,
			}

			gin.SetMode(gin.ReleaseMode)
			r := gin.Default()
			api := r.Group("/api")
			api.GET("/blogPost/:id", h.BlogDetailsHandler())

			jsonValue, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("GET", "/api/blogPost/1", bytes.NewBuffer(jsonValue))
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Result().StatusCode != tt.status {
				t.Errorf("Handler.BlogDetails() = %v, want %v", w.Result().StatusCode, tt.status)
				return
			}

			var bodyMap interface{}
			if err := json.NewDecoder(w.Body).Decode(&bodyMap); err != nil {
				t.Errorf("Handler.BlogDetails() = json body decode error %v", err)
				return
			}

			var wantBodyMap interface{}
			if enc, err := json.Marshal(tt.body); err == nil {
				_ = json.Unmarshal(enc, &wantBodyMap)
			}

			if !reflect.DeepEqual(bodyMap, wantBodyMap) {
				t.Errorf("Handler.BlogDetails() = %v, want %v", bodyMap, wantBodyMap)
			}
		})
	}
}
