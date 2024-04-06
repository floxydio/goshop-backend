package models

import "database/sql"

type User struct {
	ID       uint           `json:"id" gorm:"primaryKey;column:id"`
	Nama     string         `json:"nama" gorm:"column:nama;type:varchar(255)"`
	Username string         `json:"username" gorm:"column:username;type:varchar(50);index:unique"`
	Password string         `json:"password" gorm:"column:password;type:varchar(255)"`
	Email    sql.NullString `json:"email" gorm:"column:email;type:varchar(255)"`
	NoTelp   string         `json:"no_telp" gorm:"column:no_telp;type:varchar(20)"`
}

func (User) TableName() string {
	return "users"
}
