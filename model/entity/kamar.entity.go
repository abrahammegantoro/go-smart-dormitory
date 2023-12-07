package entity

import (
	"time"

	"github.com/google/uuid"
)

type StatusKamar string

const (
	Available StatusKamar = "Available"
	Booked    StatusKamar = "Booked"
	Occupied  StatusKamar = "Occupied"
)

type Kamar struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	NomorKamar uint8     `json:"nomor_kamar"`
	Fasilitas  string    `json:"fasilitas"`

	Status StatusKamar `json:"status"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
