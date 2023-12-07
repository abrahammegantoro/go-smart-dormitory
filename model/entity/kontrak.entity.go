package entity

import (
	"time"

	"github.com/google/uuid"
)

type Kontrak struct {
	ID            uuid.UUID `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	TanggalMasuk  time.Time `json:"tanggal_masuk"`
	TanggalKeluar time.Time `json:"tanggal_keluar"`

	// relasi dengan penghuni
	PenghuniID uuid.UUID `json:"penghuni_id" gorm:"type:uuid;index;foreignKey:ID;references:PenghuniID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	// relasi dengan kamar
	KamarID  uuid.UUID `json:"kamar_id" gorm:"type:uuid;index;foreignKey:ID;references:KamarID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" `
	PinAkses string    `json:"pin_akses"`
}
