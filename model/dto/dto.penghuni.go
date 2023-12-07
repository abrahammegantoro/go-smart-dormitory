package dto

type GetPenghuniAktifResponseDTO struct {
	NomorKamar    uint   `json:"nomor_kamar"`
	ID            uint   `json:"id"`
	Nama          string `json:"nama"`
	JenisKelamin  string `json:"jenis_kelamin"`
	NomorTelepon  string `json:"nomor_telepon"`
	KontakDarurat string `json:"kontak_darurat"`
}
