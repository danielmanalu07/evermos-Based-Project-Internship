package dtos

import (
	"evermos-app/internal/models"
	"mime/multipart"
	"time"
)

type CreateProductRequest struct {
	NamaProduk    string                  `json:"nama_produk" form:"nama_produk" validate:"required"`
	Slug          string                  `json:"slug" form:"slug" validate:"required"`
	HargaReseller string                  `json:"harga_reseller" form:"harga_reseller" validate:"required"`
	HargaKonsumen string                  `json:"harga_konsumen" form:"harga_konsumen" validate:"required"`
	Stok          int                     `json:"stok" form:"stok" validate:"required,gte=0"`
	Deskripsi     string                  `json:"deskripsi" form:"deskripsi" validate:"required"`
	IdToko        uint                    `json:"id_toko" form:"id_toko" validate:"required"`
	IdCategory    uint                    `json:"id_category" form:"id_category" validate:"required"`
	FotoProduk    []*multipart.FileHeader `json:"foto_produk" form:"foto_produk"`
}

type FotoProdukResponse struct {
	ID        uint      `json:"id"`
	IdProduk  uint      `json:"id_produk"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductResponse struct {
	ID            uint                 `json:"id"`
	NamaProduk    string               `json:"nama_produk"`
	Slug          string               `json:"slug"`
	HargaReseller string               `json:"harga_reseller"`
	HargaKonsumen string               `json:"harga_konsumen"`
	Stok          int                  `json:"stok"`
	Deskripsi     string               `json:"deskripsi"`
	IdToko        uint                 `json:"id_toko"`
	IdCategory    uint                 `json:"id_category"`
	FotoProduk    []FotoProdukResponse `json:"foto_produk"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
}

type UpdateProductRequest struct {
	NamaProduk    string `json:"nama_produk,omitempty" form:"nama_produk,omitempty"`
	Slug          string `json:"slug,omitempty" form:"slug,omitempty"`
	HargaReseller string `json:"harga_reseller,omitempty" form:"harga_reseller,omitempty"`
	HargaKonsumen string `json:"harga_konsumen,omitempty" form:"harga_konsumen,omitempty"`
	Stok          *int   `json:"stok,omitempty" form:"stok,omitempty"`
	Deskripsi     string `json:"deskripsi,omitempty" form:"deskripsi,omitempty"`
	IdCategory    *uint  `json:"id_category,omitempty" form:"id_category,omitempty"`
}

func (pr *ProductResponse) FromModel(product *models.Product) {
	pr.ID = product.ID
	pr.NamaProduk = product.NamaProduk
	pr.Slug = product.Slug
	pr.HargaReseller = product.HargaReseller
	pr.HargaKonsumen = product.HargaKonsumen
	pr.Stok = product.Stok
	pr.Deskripsi = product.Deskripsi
	pr.IdToko = product.IdToko
	pr.IdCategory = product.IdCategory
	pr.CreatedAt = product.CreatedAt
	pr.UpdatedAt = product.UpdatedAt

	for _, foto := range product.FotoProduk {
		fotoResponse := FotoProdukResponse{
			ID:        foto.ID,
			IdProduk:  foto.IdProduk,
			Url:       foto.Url,
			CreatedAt: foto.CreatedAt,
			UpdatedAt: foto.UpdatedAt,
		}
		pr.FotoProduk = append(pr.FotoProduk, fotoResponse)
	}
}

type ProductFilter struct {
	NamaProduk string `json:"nama_produk" query:"nama_produk"`
	IdToko     *uint  `json:"id_toko" query:"id_toko"`
	IdCategory *uint  `json:"id_category" query:"id_category"`
}
