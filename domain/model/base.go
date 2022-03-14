package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
Baseモデル
*/
type Base struct {
	ID        uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()" json:"id,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
