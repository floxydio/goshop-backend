package models

var userKurirTable = "users_kurir"

type UserKurir struct {
	ID             uint   `json:"id" gorm:"primaryKey;column:id"`
	Nama           string `json:"nama" gorm:"column:nama;type:varchar(255)"`
	Username       string `json:"username" gorm:"column:username;type:varchar(50);index:unique"`
	Password       string `json:"password" gorm:"column:password;type:varchar(255)"`
	KurirServiceID uint   `json:"kurir_service_id" gorm:"column:kurir_service_id"`
	LocationID     uint   `json:"location_id" gorm:"column:location_id"`
	CreatedAt      string `json:"created_at" gorm:"column:created_at"`
	KurirService   KurirService
	Location       Location
}

type UserKurirForm struct {
	Nama     string `json:"nama" form:"nama"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
	NoTelp   string `json:"no_telp" form:"no_telp"`
}

func (UserKurir) TableName() string {
	return userKurirTable
}
