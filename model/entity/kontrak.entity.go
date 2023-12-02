package entity

import "time"

type Kontrak struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	TanggalMasuk  time.Time `json:"tanggal_masuk"`
	TanggalKeluar time.Time `json:"tanggal_keluar"`

	// relasi dengan penghuni
	PenghuniID uint `json:"penghuni_id"`

	// relasi dengan kamar
	KamarID uint `json:"kamar_id"`
	PinAkses string `json:"pin_akses"`
}
