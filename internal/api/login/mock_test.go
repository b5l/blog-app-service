package login

import (
	"blog-app-service/internal/errorx"
	"blog-app-service/internal/pkg/dto"
	"context"
)

type mockDAO struct {
	login dto.LoginResponseBody

	err *errorx.Error
}

func (m *mockDAO) GetUser(ctx context.Context, username string, password string) (*dto.LoginResponseBody, *errorx.Error) {
	return &m.login, m.err
}
