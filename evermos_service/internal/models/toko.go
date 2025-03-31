package models

type Toko struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	IdUser   uint   `json:"id_user" gorm:"index;not null"`
	User     User   `json:"-" gorm:"foreignKey:IdUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	NamaToko string `json:"nama_toko" gorm:"type:varchar(255);not null"`
	UrlFoto  string `json:"url_foto" gorm:"type:varchar(255)"`
}
