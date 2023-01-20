package blogCreate

import (
	"blog-app-service/internal/errorx"
	"blog-app-service/internal/pkg/dto"
	"context"
)

type mockDAO struct {
	blogEdit dto.BlogCreateEditResponseBody

	err *errorx.Error
}

func (m *mockDAO) GetBlogCreateEdit(ctx context.Context, id int, title string, shortDescription string, longDescription string) (*dto.BlogCreateEditResponseBody, *errorx.Error) {
	return &m.blogEdit, m.err
}
