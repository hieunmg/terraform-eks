package model

import (
	"errors"
	"time"
	"weshare/common"
)

const EntityName = "Account"

type Account struct {
	common.SQLModel `json:",inline"`
	Username        string     `json:"username"`
	FullName        string     `json:"fullName"`
	Password        string     `json:"password"`
	Salt            string     `json:"salt,omitempty"`
	Avatar          string     `json:"avatar,omitempty"`
	Birthday        *time.Time `json:"birthday,omitempty"`
}

type AccountRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Salt     string `json:"salt"`
}

type AccountLoginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type AccountLoginResponse struct {
	Account
	AccessToken           string    `json:"accessToken" binding:"required"`
	RefreshToken          string    `json:"refreshToken" binding:"required"`
	RefreshTokenExpiredAt time.Time `json:"refreshTokenExpiredAt" binding:"required"`
	AccessTokenExpiredAt  time.Time `json:"accessTokenExpiredAt" binding:"required"`
}

type AccountLogoutRequest struct {
	AccountId uint32 `json:"accountId" binding:"required"`
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
