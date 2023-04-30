package service

import (
	"context"
	"database/sql"
	"time"
	"weshare/common"
	"weshare/components/hasher"
	"weshare/components/token"
	"weshare/modules/authentication/model"
	"weshare/utils"
)

type LoginRepository interface {
	Login(ctx context.Context, username string) (*model.Account, error)
	CreateSession(ctx context.Context, req *model.SessionCreate) error
}

type loginService struct {
	repository           LoginRepository
	tokenMaker           token.Maker
	bcrypt               *hasher.Bcrypt
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

func NewLoginService(
	repository LoginRepository,
	tokenMaker token.Maker,
	bcrypt *hasher.Bcrypt,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,
) *loginService {
	return &loginService{
		repository:           repository,
		tokenMaker:           tokenMaker,
		bcrypt:               bcrypt,
		accessTokenDuration:  accessTokenDuration,
		refreshTokenDuration: refreshTokenDuration,
	}
}

func (service *loginService) Login(
	ctx context.Context, req *model.AccountLoginRequest) (*model.AccountLoginResponse, error) {
	account, err := service.repository.Login(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, model.ErrUsernameOrPasswordInvalid
		}
		return nil, common.ErrDB(err)
	}

	err = service.bcrypt.CheckPassword((req.Password + account.Salt), account.Password)
	if err != nil {
		return nil, model.ErrUsernameOrPasswordInvalid
	}

	accessToken, accessPayload, err := service.tokenMaker.CreateToken(
		account.Id,
		service.accessTokenDuration)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, refreshPayload, err := service.tokenMaker.CreateToken(
		account.Id, service.refreshTokenDuration)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	err = service.repository.CreateSession(ctx, &model.SessionCreate{
		Id:           refreshPayload.Id,
		AccountId:    int(account.Id),
		RefreshToken: utils.ExtractTokenSignature(refreshToken),
		IsBlocked:    false,
		ExpiredAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return &model.AccountLoginResponse{
		Account:               *account,
		AccessToken:           accessToken,
		AccessTokenExpiredAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiredAt: refreshPayload.ExpiredAt,
	}, err
}
