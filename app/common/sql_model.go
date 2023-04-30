package common

import (
	"time"
)

type SQLModel struct {
	Id uint32 `json:"id"`
	// FakeId    *UID       `json:"id"`
	Status    int        `json:"status" default:"1"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// func (sqlModel *SQLModel) GenUID(dbType int) {
// 	uid := NewUID(uint32(sqlModel.Id), int(dbType), 1)
// 	sqlModel.FakeId = &uid
// }

// func (sqlModel *SQLModel) GetRealId() {
// 	if sqlModel.FakeId == nil {
// 		return
// 	}

// 	sqlModel.Id = uint32(sqlModel.FakeId.GetLocalID())
// }
