package model

import (
	"weshare/common"
)

type Post struct {
	common.SQLModel `json:",inline"`
	AccountId       int            `json:"account_id"`
	Description     string         `json:"description"`
	Status          int            `json:"status"`
	Images          *common.Images `json:"images,omitempty"`
}

type PostCreate struct {
	common.SQLModel `json:",inline"`
	AccountId       int            `json:"account_id"`
	Description     string         `json:"description"`
	Images          *common.Images `json:"images,omitempty"`
}
