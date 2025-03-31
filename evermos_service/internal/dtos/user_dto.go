package dtos

import "evermos-app/internal/models"

type RegisterRequest struct {
	Nama         string `json:"nama" validate:"required"`
	KataSandi    string `json:"kata_sandi" validate:"required,min=6"`
	NoTelp       string `json:"no_telp" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir" validate:"required"`
	JenisKelamin string `json:"jenis_kelamin" validate:"required"`
	Tentang      string `json:"tentang" validate:"required"`
	Pekerjaan    string `json:"pekerjaan" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	IdProvinsi   string `json:"id_provinsi" validate:"required"`
	IdKota       string `json:"id_kota" validate:"required"`
}

type LoginRequest struct {
	Email     string `json:"email" validate:"required,email"`
	KataSandi string `json:"kata_sandi" validate:"required,min=6"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type UpdateUserRequest struct {
	Nama         string `json:"nama" validate:"required"`
	KataSandi    string `json:"kata_sandi" validate:"required,min=6"`
	NoTelp       string `json:"no_telp" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir" validate:"required"`
	JenisKelamin string `json:"jenis_kelamin" validate:"required"`
	Tentang      string `json:"tentang" validate:"required"`
	Pekerjaan    string `json:"pekerjaan" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	IdProvinsi   string `json:"id_provinsi" validate:"required"`
	IdKota       string `json:"id_kota" validate:"required"`
}

type UserResponse struct {
	ID           uint   `json:"id"`
	Nama         string `json:"nama"`
	NoTelp       string `json:"notelp"`
	TanggalLahir string `json:"tanggal_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	Tentang      string `json:"tentang"`
	Pekerjaan    string `json:"pekerjaan"`
	Email        string `json:"email"`
	IdProvinsi   string `json:"id_provinsi"`
	IdKota       string `json:"id_kota"`
	IsAdmin      bool   `json:"is_admin"`
}

func (ur *UserResponse) FromModel(user *models.User) {
	ur.ID = user.ID
	ur.Nama = user.Nama
	ur.NoTelp = user.NoTelp
	ur.TanggalLahir = user.TanggalLahir.Format("02-01-2006")
	ur.JenisKelamin = user.JenisKelamin
	ur.Tentang = user.Tentang
	ur.Pekerjaan = user.Pekerjaan
	ur.Email = user.Email
	ur.IdProvinsi = user.IdProvinsi
	ur.IdKota = user.IdKota
	ur.IsAdmin = user.IsAdmin
}
