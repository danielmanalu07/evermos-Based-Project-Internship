package services

import (
	"errors"
	"evermos-app/internal/dtos"
	"evermos-app/internal/models"
	"evermos-app/internal/repositories"
)

type AlamatService struct {
	alamatRepo repositories.AlamatRepository
	userRepo   repositories.UserRepository
}

func NewAlamatService(alamatRepo repositories.AlamatRepository, userRepo repositories.UserRepository) *AlamatService {
	return &AlamatService{
		alamatRepo: alamatRepo,
		userRepo:   userRepo,
	}
}

func (s *AlamatService) CreateAlamat(userID uint, req dtos.CreateAlamatRequest) (*dtos.AlamatResponse, error) {
	if _, err := s.userRepo.FindById(userID); err != nil {
		return nil, errors.New("user not found")
	}

	alamat := &models.Alamat{
		IdUser:       userID,
		JudulAlamat:  req.JudulAlamat,
		NamaPenerima: req.NamaPenerima,
		NoTelp:       req.NoTelp,
		DetailAlamat: req.DetailAlamat,
	}

	if err := s.alamatRepo.Create(alamat); err != nil {
		return nil, err
	}

	alamatResponse := &dtos.AlamatResponse{}
	alamatResponse.FromModel(alamat)
	return alamatResponse, nil
}

func (s *AlamatService) GetAlamatByUserId(userID uint) ([]dtos.AlamatResponse, error) {
	alamats, err := s.alamatRepo.FindByUserId(userID)
	if err != nil {
		return nil, err
	}

	var alamatResponses []dtos.AlamatResponse
	for _, alamat := range alamats {
		alamatResponse := dtos.AlamatResponse{}
		alamatResponse.FromModel(&alamat)
		alamatResponses = append(alamatResponses, alamatResponse)
	}
	return alamatResponses, nil
}

func (s *AlamatService) GetAlamatById(alamatID uint) (*dtos.AlamatResponse, error) {
	alamat, err := s.alamatRepo.FindById(alamatID)
	if err != nil {
		return nil, errors.New("alamat not found")
	}

	alamatResponse := &dtos.AlamatResponse{}
	alamatResponse.FromModel(alamat)
	return alamatResponse, nil
}

func (s *AlamatService) UpdateAlamat(alamatID, userID uint, req dtos.UpdateAlamatRequest) (*dtos.AlamatResponse, error) {
	alamat, err := s.alamatRepo.FindById(alamatID)
	if err != nil {
		return nil, errors.New("alamat not found")
	}

	if alamat.IdUser != userID {
		return nil, errors.New("unauthorized: you can only update your own alamat")
	}

	if req.JudulAlamat != "" {
		alamat.JudulAlamat = req.JudulAlamat
	}
	if req.NamaPenerima != "" {
		alamat.NamaPenerima = req.NamaPenerima
	}
	if req.NoTelp != "" {
		alamat.NoTelp = req.NoTelp
	}
	if req.DetailAlamat != "" {
		alamat.DetailAlamat = req.DetailAlamat
	}

	if err := s.alamatRepo.Update(alamat); err != nil {
		return nil, err
	}

	alamatResponse := &dtos.AlamatResponse{}
	alamatResponse.FromModel(alamat)
	return alamatResponse, nil
}

func (s *AlamatService) DeleteAlamat(alamatID, userID uint) error {
	alamat, err := s.alamatRepo.FindById(alamatID)
	if err != nil {
		return errors.New("alamat not found")
	}

	if alamat.IdUser != userID {
		return errors.New("unauthorized: you can only delete your own alamat")
	}

	return s.alamatRepo.Delete(alamatID)
}
