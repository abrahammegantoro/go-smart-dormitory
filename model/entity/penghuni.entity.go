package entity

import (
	"time"

	"github.com/google/uuid"
)

type JenisKelamin string

const (
	LakiLaki  JenisKelamin = "Laki-laki"
	Perempuan JenisKelamin = "Perempuan"
)

type Status string

const (
	Diterima                 Status = "Diterima"
	MenungguPembayaran       Status = "Menunggu Pembayaran"
	MenungguPembuatanKontrak Status = "Menunggu Pembuatan Kontrak"
	BelumDireview            Status = "Belum Direview"
)

type Penghuni struct {
	ID           uuid.UUID    `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
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
