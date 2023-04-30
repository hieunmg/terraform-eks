package model

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Id           uuid.UUID `json:"id"`
	AccountId    uint32    `json:"account_id"`
	RefreshToken string    `json:"refresh_token"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiredAt    time.Time `json:"expired_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type SessionCreate struct {
	Id           uuid.UUID `json:"id"`
	AccountId    int       `json:"account_id"`
	RefreshToken string    `json:"refresh_token"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiredAt    time.Time `json:"expired_at"`
}
