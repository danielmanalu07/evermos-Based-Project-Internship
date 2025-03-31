package services

import (
	"errors"
	"evermos-app/internal/dtos"
	"evermos-app/internal/models"
	"evermos-app/internal/repositories"
	"evermos-app/pkg/auth"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo    repositories.UserRepository
	tokoRepo    repositories.TokoRepository
	wilayahRepo repositories.WilayahRepository
}

func NewUserService(userRepo repositories.UserRepository, tokoRepo repositories.TokoRepository, wilayahRepo repositories.WilayahRepository) *UserService {
	return &UserService{userRepo: userRepo, tokoRepo: tokoRepo, wilayahRepo: wilayahRepo}
}

func (s *UserService) Register(req dtos.RegisterRequest) (*models.User, error) {
	if s.userRepo.EmailExists(req.Email) {
		return nil, errors.New("email already exists")
	}

	if s.userRepo.PhoneExists(req.NoTelp) {
		return nil, errors.New("phone already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.KataSandi), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	provinsiId, err := s.wilayahRepo.GetProvinsi(req.IdProvinsi)
	if err != nil {
		return nil, err
	}

	kotaId, err := s.wilayahRepo.GetKota(req.IdKota, provinsiId)
	if err != nil {
		return nil, err
	}

	tanggalLahir, err := time.Parse("02-01-2006", req.TanggalLahir)
	if err != nil {
		return nil, errors.New("invalid tanggal_lahir format, use DD-MM-YYYY")
	}

	user := &models.User{
		Nama:         req.Nama,
		KataSandi:    string(hashedPassword),
		NoTelp:       req.NoTelp,
		TanggalLahir: tanggalLahir,
		JenisKelamin: req.JenisKelamin,
		Tentang:      req.Tentang,
		Pekerjaan:    req.Pekerjaan,
		Email:        req.Email,
		IdProvinsi:   provinsiId,
		IdKota:       kotaId,
		IsAdmin:      false,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	shop := &models.Toko{
		NamaToko: user.Nama + " Shop",
		IdUser:   user.ID,
	}

	if err := s.tokoRepo.Create(shop); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(req dtos.LoginRequest, jwtSecret string) (*dtos.LoginResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("email invalid")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.KataSandi), []byte(req.KataSandi)); err != nil {
		return nil, errors.New("password invalid")
	}

	token, err := auth.GenerateToken(user.ID, user.IsAdmin, jwtSecret)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	data := &dtos.LoginResponse{
		Token: token,
		User: dtos.UserResponse{
			ID:           user.ID,
			Nama:         user.Nama,
			NoTelp:       user.NoTelp,
			TanggalLahir: user.TanggalLahir.Format("02-01-2006"),
			JenisKelamin: user.JenisKelamin,
			Tentang:      user.Tentang,
			Pekerjaan:    user.Pekerjaan,
			Email:        user.Email,
			IdProvinsi:   user.IdProvinsi,
			IdKota:       user.IdKota,
			IsAdmin:      user.IsAdmin,
		},
	}

	return data, nil
}

func (s *UserService) GetProfile(id uint) (*dtos.UserResponse, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	userResponse := dtos.UserResponse{}
	userResponse.FromModel(user)
	return &userResponse, nil
}

func (s *UserService) UpdateProfile(userId uint, req dtos.UpdateUserRequest) (*dtos.UserResponse, error) {
	user, err := s.userRepo.FindById(userId)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if req.Nama != "" {
		user.Nama = req.Nama
	}
	if req.NoTelp != "" {
		if s.userRepo.PhoneExists(req.NoTelp) && req.NoTelp != user.NoTelp {
			return nil, errors.New("phone already exists")
		}
		user.NoTelp = req.NoTelp
	}

	if req.TanggalLahir != "" {
		tanggalLahir, err := time.Parse("02-01-2006", req.TanggalLahir)
		if err != nil {
			return nil, errors.New("invalid tanggal_lahir format, use DD-MM-YYYY")
		}
		user.TanggalLahir = tanggalLahir
	}

	if req.JenisKelamin != "" {
		user.JenisKelamin = req.JenisKelamin
	}
	if req.Tentang != "" {
		user.Tentang = req.Tentang
	}
	if req.Pekerjaan != "" {
		user.Pekerjaan = req.Pekerjaan
	}

	if req.Email != "" {
		if s.userRepo.EmailExists(req.Email) && req.Email != user.Email {
			return nil, errors.New("email already exists")
		}
		user.Email = req.Email
	}

	if req.IdProvinsi != "" {
		provinsiId, err := s.wilayahRepo.GetProvinsi(req.IdProvinsi)
		if err != nil {
			return nil, err
		}
		user.IdProvinsi = provinsiId
	}
	if req.IdKota != "" {
		kotaId, err := s.wilayahRepo.GetKota(req.IdKota, user.IdProvinsi)
		if err != nil {
			return nil, err
		}
		user.IdKota = kotaId
	}
	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}
	userResponse := dtos.UserResponse{}
	userResponse.FromModel(user)
	return &userResponse, nil
}

func (s *UserService) Logout(token string) error {
	if s.userRepo.IsTokenRevoked(token) {
		return errors.New("token already revoked")
	}
	return s.userRepo.RevokeToken(token)
}

func (s *UserService) DeleteAccount(userID uint, token string) error {
	_, err := s.userRepo.FindById(userID)
	if err != nil {
		return errors.New("user not found")
	}

	if err := s.userRepo.Delete(userID); err != nil {
		return err
	}

	if token != "" {
		if err := s.userRepo.RevokeToken(token); err != nil {
			return err
		}
	}

	return nil
}
