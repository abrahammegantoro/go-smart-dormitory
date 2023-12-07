package entity

import (
	"time"

	"github.com/google/uuid"
)

type Kontrak struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	TanggalMasuk  time.Time `json:"tanggal_masuk"`
	TanggalKeluar time.Time `json:"tanggal_keluar"`

	// relasi dengan penghuni
	PenghuniID uint `json:"penghuni_id"`

	// relasi dengan kamar
	KamarID  uint   `json:"kamar_id"`
	PinAkses string `json:"pin_akses"`
}
