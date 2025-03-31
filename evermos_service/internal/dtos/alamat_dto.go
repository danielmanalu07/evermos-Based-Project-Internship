package dtos

import (
	"evermos-app/internal/models"
	"time"
)

type CreateAlamatRequest struct {
	JudulAlamat  string `json:"judul_alamat" validate:"required"`
	NamaPenerima string `json:"nama_penerima" validate:"required"`
	NoTelp       string `json:"no_telp" validate:"required"`
	DetailAlamat string `json:"detail_alamat" validate:"required"`
}

type UpdateAlamatRequest struct {
	JudulAlamat  string `json:"judul_alamat,omitempty"`
	NamaPenerima string `json:"nama_penerima,omitempty"`
	NoTelp       string `json:"no_telp,omitempty"`
	DetailAlamat string `json:"detail_alamat,omitempty"`
}

type AlamatResponse struct {
	ID           uint      `json:"id"`
	IdUser       uint      `json:"id_user"`
	JudulAlamat  string    `json:"judul_alamat"`
	NamaPenerima string    `json:"nama_penerima"`
	NoTelp       string    `json:"no_telp"`
	DetailAlamat string    `json:"detail_alamat"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (ar *AlamatResponse) FromModel(alamat *models.Alamat) {
	ar.ID = alamat.ID
	ar.IdUser = alamat.IdUser
	ar.JudulAlamat = alamat.JudulAlamat
	ar.NamaPenerima = alamat.NamaPenerima
	ar.NoTelp = alamat.NoTelp
	ar.DetailAlamat = alamat.DetailAlamat
	ar.CreatedAt = alamat.CreatedAt
	ar.UpdatedAt = alamat.UpdatedAt
}
