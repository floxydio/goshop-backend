package models

var tableName = "kurir_service"

type KurirService struct {
	KurirServiceID uint   `json:"id" gorm:"primaryKey;column:kurir_service_id"`
	JenisLayanan   string `json:"jenis_layanan" gorm:"column:jenis_layanan;type:varchar(255)"`
	Status         string `json:"status" gorm:"column:status;type:int(11);default:0"`
}

func (KurirService) TableName() string {
	return tableName
}
