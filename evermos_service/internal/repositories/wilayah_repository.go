package repositories

type WilayahRepository interface {
	GetProvinsi(namaProvinsi string) (string, error)
	GetKota(namaKota, provinsi_id string) (string, error)
}
