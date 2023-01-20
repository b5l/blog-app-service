package blogDetails

import (
	"blog-app-service/internal/errorx"
	"blog-app-service/internal/pkg/dto"
	"context"
)

type mockDAO struct {
	blogDetails dto.BlogDetailsResponseBody

	err *errorx.Error
}

func (m *mockDAO) GetBlogDetails(ctx context.Context, id string) (dto.BlogDetailsResponseBody, *errorx.Error) {
	return m.blogDetails, m.err
}
