package appcontext

import (
	"database/sql"
	"time"
	"weshare/components/uploadprovider"
)

type AppContext interface {
	GetDBConnection() *sql.DB
	GetSecretKey() string
	GetAccessTokenDuration() time.Duration
	GetRefreshTokenDuration() time.Duration
	GetUploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db              *sql.DB
	secretKey       string
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	s3Provider      uploadprovider.UploadProvider
}

func NewAppContext(
	db *sql.DB,
	secretKey string,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration,
	s3Provider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{
		db:              db,
		secretKey:       secretKey,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
		s3Provider:      s3Provider,
	}
}

func (ctx *appCtx) GetDBConnection() *sql.DB {
	return ctx.db
}

func (ctx *appCtx) GetAccessTokenDuration() time.Duration {
	return ctx.accessTokenTTL
}

func (ctx *appCtx) GetRefreshTokenDuration() time.Duration {
	return ctx.refreshTokenTTL
}

func (ctx *appCtx) GetSecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetUploadProvider() uploadprovider.UploadProvider {
	return ctx.s3Provider
}
