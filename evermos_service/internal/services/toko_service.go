package services

import (
	"errors"
	"evermos-app/internal/dtos"
	"evermos-app/internal/repositories"
	"evermos-app/pkg/storage"
)

type TokoService struct {
	tokoRepo repositories.TokoRepository
	userRepo repositories.UserRepository
}

func NewTokoService(tokoRepo repositories.TokoRepository, userRepo repositories.UserRepository) *TokoService {
	return &TokoService{
		tokoRepo: tokoRepo,
		userRepo: userRepo,
	}
}

func (s *TokoService) GetTokoByUserId(userID uint) (*dtos.TokoResponse, error) {
	tokos, err := s.tokoRepo.FindByUserId(userID)
	if err != nil {
		return nil, err
	}

	var tokoResponse dtos.TokoResponse
	tokoResponse.FromModel(tokos)
	return &tokoResponse, nil
}

func (s *TokoService) GetTokoById(tokoID uint) (*dtos.TokoResponse, error) {
	toko, err := s.tokoRepo.FindById(tokoID)
	if err != nil {
		return nil, errors.New("toko not found")
	}

	tokoResponse := &dtos.TokoResponse{}
	tokoResponse.FromModel(toko)
	return tokoResponse, nil
}

func (s *TokoService) UpdateToko(tokoID, userID uint, req dtos.UpdateTokoRequest) (*dtos.TokoResponse, error) {
	toko, err := s.tokoRepo.FindById(tokoID)
	if err != nil {
		return nil, errors.New("toko not found")
	}

	// Validasi kepemilikan
	if toko.IdUser != userID {
		return nil, errors.New("unauthorized: you can only update your own toko")
	}

	if req.NamaToko != "" {
		toko.NamaToko = req.NamaToko
	}
	if req.UrlFoto != nil {
		filePath, err := storage.SaveFile(req.UrlFoto)
		if err != nil {
			return nil, errors.New("failed to upload file")
		}
		toko.UrlFoto = filePath
	}

	if err := s.tokoRepo.Update(toko); err != nil {
		return nil, err
	}

	tokoResponse := &dtos.TokoResponse{}
	tokoResponse.FromModel(toko)
	return tokoResponse, nil
}

func (s *TokoService) DeleteToko(tokoID, userID uint) error {
	toko, err := s.tokoRepo.FindById(tokoID)
	if err != nil {
		return errors.New("toko not found")
	}

	if toko.IdUser != userID {
		return errors.New("unauthorized: you can only delete your own toko")
	}

	if toko.UrlFoto != "" {
		err := storage.DeleteFile(toko.UrlFoto)
		if err != nil {
			return errors.New("failed to delete image file")
		}
	}

	return s.tokoRepo.Delete(tokoID)
}
