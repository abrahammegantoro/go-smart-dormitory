package entity

import (
	"time"

	"github.com/google/uuid"
)

type Admin struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username  string    `json:"username"`
	Password  string    `json:"-" gorm:"column:password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
