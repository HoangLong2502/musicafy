package common

import "time"

type SQLModel struct {
	ID        int        `json:"_" gorm:"primaryKey;column:id"`
	FakeId    *UID       `json:"id"`
	Status    int        `json:"status"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (m *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.ID), dbType, 1)
	m.FakeId = &uid
}
