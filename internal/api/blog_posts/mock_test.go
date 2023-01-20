package blogPosts

import (
	"blog-app-service/internal/errorx"
	"blog-app-service/internal/pkg/dto"
	"context"
)

type mockDAO struct {
	blogPosts []dto.BlogPostsObject

	err *errorx.Error
}

func (m *mockDAO) GetBlogPosts(ctx context.Context) ([]dto.BlogPostsObject, *errorx.Error) {
	return m.blogPosts, m.err
}
