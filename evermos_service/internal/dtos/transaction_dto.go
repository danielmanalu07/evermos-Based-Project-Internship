package dtos

import (
	"evermos-app/internal/models"
	"time"
)

type CreateTransaksiRequest struct {
	AlamatPengiriman uint                  `json:"alamat_pengiriman" validate:"required"`
	MethodBayar      string                `json:"method_bayar" validate:"required"`
	Items            []CreateTransaksiItem `json:"items" validate:"required,dive"`
}
type CreateTransaksiItem struct {
	ProductID uint `json:"product_id" validate:"required"`
	Kuantitas int  `json:"kuantitas" validate:"required,gt=0"`
}

type LogProdukResponse struct {
	ID            uint      `json:"id"`
	IdProduk      int       `json:"id_produk"`
	NamaProduk    string    `json:"nama_produk"`
	Slug          string    `json:"slug"`
	HargaReseller string    `json:"harga_reseller"`
	HargaKonsumen string    `json:"harga_konsumen"`
	Deskripsi     string    `json:"deskripsi"`
	IdToko        uint      `json:"id_toko"`
	IdCategory    uint      `json:"id_category"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type DetailTransaksiResponse struct {
	ID          uint              `json:"id"`
	IdTrx       uint              `json:"id_trx"`
	IdLogProduk uint              `json:"id_log_produk"`
	LogProduk   LogProdukResponse `json:"log_produk"`
	IdToko      uint              `json:"id_toko"`
	Kuantitas   int               `json:"kuantitas"`
	HargaTotal  int64             `json:"harga_total"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

type TransaksiResponse struct {
	ID               uint                      `json:"id"`
	IdUser           uint                      `json:"id_user"`
	AlamatPengiriman uint                      `json:"alamat_pengiriman"`
	HargaTotal       int64                     `json:"harga_total"`
	KodeInvoice      string                    `json:"kode_invoice"`
	MethodBayar      string                    `json:"method_bayar"`
	Details          []DetailTransaksiResponse `json:"details"`
	CreatedAt        time.Time                 `json:"created_at"`
	UpdatedAt        time.Time                 `json:"updated_at"`
}

func (tr *TransaksiResponse) FromModel(transaksi *models.Transaksi, details []models.DetailTransaksi) {
	tr.ID = transaksi.ID
	tr.IdUser = transaksi.IdUser
	tr.AlamatPengiriman = transaksi.AlamatPengiriman
	tr.HargaTotal = transaksi.HargaTotal
	tr.KodeInvoice = transaksi.KodeInvoice
	tr.MethodBayar = transaksi.MethodBayar
	tr.CreatedAt = transaksi.CreatedAt
	tr.UpdatedAt = transaksi.UpdatedAt

	for _, detail := range details {
		logProdukResp := LogProdukResponse{
			ID:            detail.LogProduk.ID,
			IdProduk:      detail.LogProduk.IdProduk,
			NamaProduk:    detail.LogProduk.NamaProduk,
			Slug:          detail.LogProduk.Slug,
			HargaReseller: detail.LogProduk.HargaReseller,
			HargaKonsumen: detail.LogProduk.HargaKonsumen,
			Deskripsi:     detail.LogProduk.Deskripsi,
			IdToko:        detail.LogProduk.IdToko,
			IdCategory:    detail.LogProduk.IdCategory,
			CreatedAt:     detail.LogProduk.CreatedAt,
			UpdatedAt:     detail.LogProduk.UpdatedAt,
		}
		detailResp := DetailTransaksiResponse{
			ID:          detail.ID,
			IdTrx:       detail.IdTrx,
			IdLogProduk: detail.IdLogProduk,
			LogProduk:   logProdukResp,
			IdToko:      detail.IdToko,
			Kuantitas:   detail.Kuantitas,
			HargaTotal:  detail.HargaTotal,
			CreatedAt:   detail.CreatedAt,
			UpdatedAt:   detail.UpdatedAt,
		}
		tr.Details = append(tr.Details, detailResp)
	}
}

type TransactionFilter struct {
	KodeInvoice string `json:"kode_invoice" query:"kode_invoice"`
	MethodBayar string `json:"method_bayar" query:"method_bayar"`
}
