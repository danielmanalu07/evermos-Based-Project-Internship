package services

import (
	"errors"
	"evermos-app/internal/dtos"
	"evermos-app/internal/models"
	"evermos-app/internal/repositories"
	"fmt"
	"strconv"
	"time"
)

type TransactionService struct {
	transRepo   repositories.TransactionRepository
	userRepo    repositories.UserRepository
	alamatRepo  repositories.AlamatRepository
	productRepo repositories.ProductRepository
	tokoRepo    repositories.TokoRepository
}

func NewTransactionService(transRepo repositories.TransactionRepository, userRepo repositories.UserRepository, alamatRepo repositories.AlamatRepository, productRepo repositories.ProductRepository, tokoRepo repositories.TokoRepository) *TransactionService {
	return &TransactionService{
		transRepo:   transRepo,
		userRepo:    userRepo,
		alamatRepo:  alamatRepo,
		productRepo: productRepo,
		tokoRepo:    tokoRepo,
	}
}

func (s *TransactionService) CreateTransaksi(userID uint, req dtos.CreateTransaksiRequest) (*dtos.TransaksiResponse, error) {
	alamat, err := s.alamatRepo.FindById(req.AlamatPengiriman)
	if err != nil {
		return nil, errors.New("alamat not found")
	}
	if alamat.IdUser != userID {
		return nil, errors.New("unauthorized: you can only use your own alamat")
	}

	var totalHarga int64
	for _, item := range req.Items {
		product, err := s.productRepo.FindById(item.ProductID)
		if err != nil {
			return nil, errors.New("product not found: " + strconv.Itoa(int(item.ProductID)))
		}
		if product.Stok < item.Kuantitas {
			return nil, errors.New("insufficient stock for product: " + product.NamaProduk)
		}
		hargaKonsumen, err := strconv.ParseInt(product.HargaKonsumen, 10, 64)
		if err != nil {
			return nil, errors.New("invalid harga_konsumen for product: " + product.NamaProduk)
		}
		totalHarga += hargaKonsumen * int64(item.Kuantitas)
	}

	kodeInvoice := fmt.Sprintf("TRX-%d", time.Now().Unix())

	transaksi := &models.Transaksi{
		IdUser:           userID,
		AlamatPengiriman: req.AlamatPengiriman,
		HargaTotal:       totalHarga,
		KodeInvoice:      kodeInvoice,
		MethodBayar:      req.MethodBayar,
	}
	if err := s.transRepo.CreateTransaksi(transaksi); err != nil {
		return nil, err
	}

	for _, item := range req.Items {
		product, _ := s.productRepo.FindById(item.ProductID)

		logProduk := &models.LogProduk{
			IdProduk:      int(product.ID),
			NamaProduk:    product.NamaProduk,
			Slug:          product.Slug,
			HargaReseller: product.HargaReseller,
			HargaKonsumen: product.HargaKonsumen,
			Deskripsi:     product.Deskripsi,
			IdToko:        product.IdToko,
			IdCategory:    product.IdCategory,
		}
		logProdukCreated, err := s.transRepo.CreateLogProduk(logProduk)
		if err != nil {
			return nil, errors.New("failed to create log produk: " + err.Error())
		}

		product.Stok -= item.Kuantitas
		if err := s.productRepo.Update(product); err != nil {
			return nil, errors.New("failed to update product stock: " + err.Error())
		}

		hargaKonsumen, _ := strconv.ParseInt(product.HargaKonsumen, 10, 64)
		itemTotal := hargaKonsumen * int64(item.Kuantitas)

		detail := &models.DetailTransaksi{
			IdTrx:       transaksi.ID,
			IdLogProduk: logProdukCreated.ID,
			IdToko:      product.IdToko,
			Kuantitas:   item.Kuantitas,
			HargaTotal:  itemTotal,
		}
		if err := s.transRepo.CreateDetailTransaksi(detail); err != nil {
			return nil, errors.New("failed to create detail transaksi: " + err.Error())
		}
	}

	createdTransaksi, err := s.transRepo.FindTransaksiById(transaksi.ID)
	if err != nil {
		return nil, err
	}
	details, err := s.transRepo.FindDetailsByTransaksiId(transaksi.ID)
	if err != nil {
		return nil, err
	}

	transaksiResp := &dtos.TransaksiResponse{}
	transaksiResp.FromModel(createdTransaksi, details)
	return transaksiResp, nil
}

func (s *TransactionService) GetTransaksiById(transaksiID, userID uint) (*dtos.TransaksiResponse, error) {
	transaksi, err := s.transRepo.FindTransaksiById(transaksiID)
	if err != nil {
		return nil, errors.New("transaksi not found")
	}

	if transaksi.IdUser != userID {
		return nil, errors.New("unauthorized: you can only view your own transaksi")
	}

	details, err := s.transRepo.FindDetailsByTransaksiId(transaksiID)
	if err != nil {
		return nil, err
	}

	transaksiResp := &dtos.TransaksiResponse{}
	transaksiResp.FromModel(transaksi, details)
	return transaksiResp, nil
}

func (s *TransactionService) GetAllTransaksis(userID uint) ([]dtos.TransaksiResponse, error) {
	transaksis, err := s.transRepo.FindTransaksisByUserId(userID)
	if err != nil {
		return nil, errors.New("no transaksis found")
	}

	var transaksiResponses []dtos.TransaksiResponse
	for _, transaksi := range transaksis {
		details, err := s.transRepo.FindDetailsByTransaksiId(transaksi.ID)
		if err != nil {
			continue
		}
		transaksiResp := dtos.TransaksiResponse{}
		transaksiResp.FromModel(&transaksi, details)
		transaksiResponses = append(transaksiResponses, transaksiResp)
	}
	if len(transaksiResponses) == 0 {
		return nil, errors.New("no transaksis found")
	}
	return transaksiResponses, nil
}

func (s *TransactionService) GetPageAndFilterTransaksis(userID uint, page, limit int, filter *dtos.TransactionFilter) (*dtos.PaginatedResponse, error) {
	transaksis, total, err := s.transRepo.FindTransaksisWithPaginationAndFilter(userID, page, limit, (*models.TransactionFilter)(filter))
	if err != nil {
		return nil, errors.New("no transaksis found")
	}

	var transaksiResponses []dtos.TransaksiResponse
	for _, transaksi := range transaksis {
		details, err := s.transRepo.FindDetailsByTransaksiId(transaksi.ID)
		if err != nil {
			continue
		}
		transaksiResp := dtos.TransaksiResponse{}
		transaksiResp.FromModel(&transaksi, details)
		transaksiResponses = append(transaksiResponses, transaksiResp)
	}

	totalPages := (int(total) + limit - 1) / limit
	meta := dtos.PaginationMeta{
		TotalItems:  int(total),
		TotalPages:  totalPages,
		CurrentPage: page,
		Limit:       limit,
	}

	return &dtos.PaginatedResponse{
		Data: transaksiResponses,
		Meta: meta,
	}, nil
}
