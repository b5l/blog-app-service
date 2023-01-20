package blogDelete

import (
	"blog-app-service/internal/errorx"
	"blog-app-service/internal/pkg/dto"
	"context"
)

type mockDAO struct {
	blogDelete dto.BlogDeleteResponseBody

	err *errorx.Error
}

func (m *mockDAO) GetBlogDelete(ctx context.Context, id string) (dto.BlogDeleteResponseBody, *errorx.Error) {
	return m.blogDelete, m.err
}
