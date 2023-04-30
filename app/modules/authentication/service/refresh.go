package service

import (
	"context"
	"errors"
	"time"
	"weshare/common"
	"weshare/components/token"
	"weshare/modules/authentication/model"
	"weshare/utils"

	"github.com/google/uuid"
)

type renewAccessTokenRepository interface {
	GetAccountById(ctx context.Context, id uint32) (*model.Account, error)
	GetSession(ctx context.Context, id uuid.UUID) (*model.Session, error)
}

type renewAccessTokenService struct {
	repository       renewAccessTokenRepository
	tokenMaker       token.Maker
	accessTokenTimer time.Duration
}

func NewRenewAccessTokenService(repository renewAccessTokenRepository, tokenMaker token.Maker, accessTokenTimer time.Duration,
) *renewAccessTokenService {
	return &renewAccessTokenService{
		repository:       repository,
		tokenMaker:       tokenMaker,
		accessTokenTimer: accessTokenTimer,
	}
}

func (service *renewAccessTokenService) RenewAccessToken(ctx context.Context, req *model.RenewAccessTokenRequest) (*model.RenewAccessTokenResponse, error) {

	refreshPayload, err := service.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return nil, common.ErrUnAuthorized(err)
	}

	account, err := service.repository.GetAccountById(ctx, uint32(refreshPayload.AccountId))
	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrEntityNotFound(model.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}

	session, err := service.repository.GetSession(ctx, refreshPayload.Id)
	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrEntityNotFound("Session", err)
		}
		return nil, common.ErrDB(err)
	}

	if session.AccountId != refreshPayload.AccountId {
		message := "incorrect session user"
		panic(common.NewUnauthorized(errors.New(message), message, "Unauthorized"))
	}

	if session.RefreshToken != utils.ExtractTokenSignature(req.RefreshToken) {
		message := "mismatched session token"
		panic(common.NewUnauthorized(errors.New(message), message, "Unauthorized"))
	}

	if time.Now().After(session.ExpiredAt) {
		message := "expired session"
		panic(common.NewUnauthorized(errors.New(message), message, "Unauthorized"))
	}

	accessToken, accessPayload, err := service.tokenMaker.CreateToken(
		account.Id,
		service.accessTokenTimer,
	)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return &model.RenewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiredAt: accessPayload.ExpiredAt,
	}, nil
}
