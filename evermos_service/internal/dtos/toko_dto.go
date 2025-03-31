package dtos

import (
	"evermos-app/internal/models"
	"mime/multipart"
)

type CreateTokoRequest struct {
	NamaToko string                `json:"nama_toko" validate:"required"`
	UrlFoto  *multipart.FileHeader `json:"url_foto,omitempty"`
}

type UpdateTokoRequest struct {
	NamaToko string                `json:"nama_toko,omitempty"`
	UrlFoto  *multipart.FileHeader `json:"url_foto,omitempty"`
}

type TokoResponse struct {
	ID       uint   `json:"id"`
	IdUser   uint   `json:"id_user"`
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
}

func (tr *TokoResponse) FromModel(toko *models.Toko) {
	tr.ID = toko.ID
	tr.IdUser = toko.IdUser
	tr.NamaToko = toko.NamaToko
	tr.UrlFoto = toko.UrlFoto
}
