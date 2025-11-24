package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

func (t *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().Unix()
	t.CreatedAt = now
	t.UpdatedAt = now
	return
}

func (t *Todo) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now().Unix()
	t.UpdatedAt = now
	return
}
