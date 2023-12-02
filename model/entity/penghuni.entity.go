package entity

import (
	"time"
)

type JenisKelamin string

const (
	LakiLaki  JenisKelamin = "Laki-laki"
	Perempuan JenisKelamin = "Perempuan"
)

type Status string

const (
	Diterima                 Status = "Diterima"
	MenungguAlokasiKamar     Status = "Menunggu Alokasi Kamar"
	MenungguPembayaran       Status = "Menunggu Pembayaran"
	MenungguPembuatanKontrak Status = "Menunggu Pembuatan Kontrak"
	BelumDireview            Status = "Belum Direview"
)

type Penghuni struct {
	ID           uint         `json:"id" gorm:"primaryKey"`
	Email        string       `json:"email" gorm:"unique"`
	Nama         string       `json:"nama"`
	NIM          string       `json:"nim"`
	JenisKelamin JenisKelamin `json:"jenis_kelamin"`
	NomorTelepon string       `json:"nomor_telepon"`

	// nama dan hubungan kontak darurat
	KontakDarurat         string `json:"kontak_darurat"`
	NamaKontakDarurat     string `json:"nama_kontak_darurat"`
	HubunganKontakDarurat string `json:"hubungan_kontak_darurat"`

	Alasan string `json:"alasan" gorm:"type:text"`
	Status Status `json:"status"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
