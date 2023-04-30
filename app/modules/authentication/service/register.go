package service

import (
	"context"
	"weshare/common"
	"weshare/components/hasher"
	"weshare/modules/authentication/model"
)

type registerRepository interface {
	Register(ctx context.Context, req *model.AccountRegisterRequest) error
}

type registerService struct {
	repository registerRepository
	bcrypt     *hasher.Bcrypt
}

func NewRegisterService(repository registerRepository, bcrypt *hasher.Bcrypt) *registerService {
	return &registerService{
		repository: repository,
		bcrypt:     bcrypt,
	}
}

func (service *registerService) RegisterAccount(
	ctx context.Context,
	req *model.AccountRegisterRequest) error {

	salt := hasher.GenSalt(6)
	hashedPassword, err := service.bcrypt.HashPassword(req.Password + salt)

	if err != nil {
		return common.ErrInternal(err)
	}

	err = service.repository.Register(ctx, &model.AccountRegisterRequest{
		Username: req.Username,
		FullName: req.FullName,
		Password: hashedPassword,
		Salt:     salt,
	})

	if err != nil {
		return common.ErrDB(err)
	}

	return nil
}
