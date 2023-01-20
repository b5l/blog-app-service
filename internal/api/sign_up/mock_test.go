package signUp

import (
	"blog-app-service/internal/errorx"
	"blog-app-service/internal/pkg/dto"
	"context"
)

type mockDAO struct {
	signUp dto.SignUpResponseBody

	err *errorx.Error
}

func (m *mockDAO) PostUser(ctx context.Context, username string, password string) (dto.SignUpResponseBody, *errorx.Error) {
	return m.signUp, m.err
}
