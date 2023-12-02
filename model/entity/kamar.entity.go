package entity

import (
	"time"
)

type StatusKamar string

const (
	Available StatusKamar = "Available"
	Booked    StatusKamar = "Booked"
	Occupied  StatusKamar = "Occupied"
)

type Kamar struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	NomorKamar uint8  `json:"nomor_kamar"`
	Fasilitas  string `json:"fasilitas"`

	Status StatusKamar `json:"status"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
