package repositoryimpl

import (
	"encoding/json"
	"errors"
	"evermos-app/internal/models"
	"evermos-app/internal/repositories"
	"io"
	"net/http"

	"gorm.io/gorm"
)

type wilayahRepositoryImpl struct {
	db *gorm.DB
}

func NewWilayahRepository(db *gorm.DB) repositories.WilayahRepository {
	return &wilayahRepositoryImpl{db: db}
}

// GetProvinsi mengambil ID provinsi berdasarkan nama
func (w *wilayahRepositoryImpl) GetProvinsi(idProvinsi string) (string, error) {
	resp, err := http.Get("https://emsifa.github.io/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var provinsis []models.Provinsi
	if err := json.Unmarshal(body, &provinsis); err != nil {
		return "", err
	}

	for _, provinsi := range provinsis {
		if provinsi.ID == idProvinsi {
			return provinsi.ID, nil
		}
	}
	return "", errors.New("provinsi tidak ditemukan: " + idProvinsi)
}

// GetKota mengambil ID kota berdasarkan nama dan ID provinsi
func (w *wilayahRepositoryImpl) GetKota(idKota string, provinsiID string) (string, error) {
	resp, err := http.Get("https://emsifa.github.io/api-wilayah-indonesia/api/regencies/" + provinsiID + ".json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var kotas []models.Kota
	if err := json.Unmarshal(body, &kotas); err != nil {
		return "", err
	}

	for _, kota := range kotas {
		if kota.ID == idKota {
			return kota.ID, nil
		}
	}
	return "", errors.New("kota tidak ditemukan: " + idKota + " di provinsi " + provinsiID)
}
