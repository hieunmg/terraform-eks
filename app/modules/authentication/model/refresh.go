package model

import (
	"time"
)

type RenewAccessTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
} // @name RenewAccessTokenRequest

type RenewAccessTokenResponse struct {
	AccessToken          string    `json:"accessToken" binding:"required"`
	AccessTokenExpiredAt time.Time `json:"accessTokenExpiredAt" binding:"required"`
} // @name RenewAccessTokenResponse
