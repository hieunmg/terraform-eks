package service

import (
	"context"
)

type LogoutRepository interface {
	Logout(context.Context, uint32) error
}

type logoutService struct {
	repository LogoutRepository
}

func NewLogoutRepository(repository LogoutRepository) *logoutService {
	return &logoutService{
		repository: repository,
	}
}

func (service *logoutService) Logout(ctx context.Context, accountId uint32) error {
	err := service.repository.Logout(ctx, accountId)
	return err
}
