package services

import (
	"errors"
	"evermos-app/internal/dtos"
	"evermos-app/internal/models"
	"evermos-app/internal/repositories"
	"evermos-app/pkg/storage"
	"mime/multipart"
)

type ProductService struct {
	productRepo  repositories.ProductRepository
	tokoRepo     repositories.TokoRepository
	categoryRepo repositories.CategoryRepository
}

func NewProductService(productRepo repositories.ProductRepository, tokoRepo repositories.TokoRepository, categoryRepo repositories.CategoryRepository) *ProductService {
	return &ProductService{
		productRepo:  productRepo,
		tokoRepo:     tokoRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *ProductService) CreateProduct(userID uint, req dtos.CreateProductRequest, files []*multipart.FileHeader) (*dtos.ProductResponse, error) {
	toko, err := s.tokoRepo.FindById(req.IdToko)
	if err != nil {
		return nil, errors.New("toko not found")
	}
	if toko.IdUser != userID {
		return nil, errors.New("unauthorized: you can only add products to your own toko")
	}

	if _, err := s.categoryRepo.FindById(req.IdCategory); err != nil {
		return nil, errors.New("category not found")
	}

	product := &models.Product{
		NamaProduk:    req.NamaProduk,
		Slug:          req.Slug,
		HargaReseller: req.HargaReseller,
		HargaKonsumen: req.HargaKonsumen,
		Stok:          req.Stok,
		Deskripsi:     req.Deskripsi,
		IdToko:        req.IdToko,
		IdCategory:    req.IdCategory,
	}

	if err := s.productRepo.Create(product); err != nil {
		return nil, err
	}

	for _, file := range files {
		filePath, err := storage.SaveFile(file)
		if err != nil {
			return nil, errors.New("failed to upload foto produk: " + err.Error())
		}

		fotoProduk := &models.FotoProduk{
			IdProduk: product.ID,
			Url:      filePath,
		}
		if err := s.productRepo.CreateFotoProduk(fotoProduk); err != nil {
			return nil, errors.New("failed to save foto produk: " + err.Error())
		}
	}

	productWithFotos, err := s.productRepo.FindByIdWithFotos(product.ID)
	if err != nil {
		return nil, errors.New("failed to fetch product with fotos: " + err.Error())
	}

	productResponse := &dtos.ProductResponse{}
	productResponse.FromModel(productWithFotos)

	return productResponse, nil

}

func (s *ProductService) GetAllProducts(userId uint) ([]dtos.ProductResponse, error) {
	products, err := s.productRepo.FindAll()
	if err != nil {
		return nil, err
	}

	// Filter hanya produk dari toko milik pengguna
	var productResponses []dtos.ProductResponse
	for _, product := range products {
		toko, err := s.tokoRepo.FindById(product.IdToko)
		if err == nil && toko.IdUser == userId {
			productResponse := dtos.ProductResponse{}
			productResponse.FromModel(&product)
			productResponses = append(productResponses, productResponse)
		}
	}
	if len(productResponses) == 0 {
		return nil, errors.New("no products found for your toko")
	}
	return productResponses, nil
}

func (s *ProductService) GetProductsByTokoId(tokoID, userID uint) ([]dtos.ProductResponse, error) {
	toko, err := s.tokoRepo.FindById(tokoID)
	if err != nil {
		return nil, errors.New("toko not found")
	}
	if toko.IdUser != userID {
		return nil, errors.New("unauthorized: you can only view products from your own toko")
	}

	products, err := s.productRepo.FindByTokoId(tokoID)
	if err != nil {
		return nil, err
	}

	var productResponses []dtos.ProductResponse
	for _, product := range products {
		productResponse := dtos.ProductResponse{}
		productResponse.FromModel(&product)
		productResponses = append(productResponses, productResponse)
	}
	return productResponses, nil
}

func (s *ProductService) GetProductById(productID, userID uint) (*dtos.ProductResponse, error) {
	product, err := s.productRepo.FindById(productID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	toko, err := s.tokoRepo.FindById(product.IdToko)
	if err != nil {
		return nil, errors.New("toko not found")
	}
	if toko.IdUser != userID {
		return nil, errors.New("unauthorized: you can only view products from your own toko")
	}

	productResponse := &dtos.ProductResponse{}
	productResponse.FromModel(product)
	return productResponse, nil
}

func (s *ProductService) UpdateProduct(productID, userID uint, req dtos.UpdateProductRequest) (*dtos.ProductResponse, error) {
	product, err := s.productRepo.FindById(productID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	toko, err := s.tokoRepo.FindById(product.IdToko)
	if err != nil {
		return nil, errors.New("toko not found")
	}
	if toko.IdUser != userID {
		return nil, errors.New("unauthorized: you can only update products in your own toko")
	}

	if req.NamaProduk != "" {
		product.NamaProduk = req.NamaProduk
	}
	if req.Slug != "" {
		products, _ := s.productRepo.FindAll()
		for _, p := range products {
			if p.Slug == req.Slug && p.ID != productID {
				return nil, errors.New("slug already exists")
			}
		}
		product.Slug = req.Slug
	}
	if req.HargaReseller != "" {
		product.HargaReseller = req.HargaReseller
	}
	if req.HargaKonsumen != "" {
		product.HargaKonsumen = req.HargaKonsumen
	}
	if req.Stok != nil {
		if *req.Stok < 0 {
			return nil, errors.New("stock cannot be negative")
		}
		product.Stok = *req.Stok
	}
	if req.Deskripsi != "" {
		product.Deskripsi = req.Deskripsi
	}
	if req.IdCategory != nil {
		if _, err := s.categoryRepo.FindById(*req.IdCategory); err != nil {
			return nil, errors.New("category not found")
		}
		product.IdCategory = *req.IdCategory
	}

	if err := s.productRepo.Update(product); err != nil {
		return nil, err
	}

	productResponse := &dtos.ProductResponse{}
	productResponse.FromModel(product)
	return productResponse, nil
}

func (s *ProductService) DeleteProduct(productID, userID uint) error {
	product, err := s.productRepo.FindById(productID)
	if err != nil {
		return errors.New("product not found")
	}

	toko, err := s.tokoRepo.FindById(product.IdToko)
	if err != nil {
		return errors.New("toko not found")
	}
	if toko.IdUser != userID {
		return errors.New("unauthorized: you can only delete products in your own toko")
	}

	return s.productRepo.Delete(productID)
}

func (s *ProductService) GetPageAndFilterProducts(userID uint, page, limit int, filter *dtos.ProductFilter) (*dtos.PaginatedResponse, error) {
	products, total, err := s.productRepo.FindAllWithPaginationAndFilter(userID, page, limit, filter)
	if err != nil {
		return nil, errors.New("no products found matching the criteria")
	}

	var productResponses []dtos.ProductResponse
	for _, product := range products {
		productResponse := dtos.ProductResponse{}
		productResponse.FromModel(&product)
		productResponses = append(productResponses, productResponse)
	}

	totalPages := (int(total) + limit - 1) / limit
	meta := dtos.PaginationMeta{
		TotalItems:  int(total),
		TotalPages:  totalPages,
		CurrentPage: page,
		Limit:       limit,
	}

	return &dtos.PaginatedResponse{
		Data: productResponses,
		Meta: meta,
	}, nil
}
