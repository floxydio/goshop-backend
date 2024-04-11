package models

var komplainTable = "komplain"

type Komplain struct {
	KomplainID uint   `json:"id" gorm:"primaryKey;column:komplain_id"`
	UserID     uint   `json:"user_id" gorm:"column:user_id"`
	ProductID  uint   `json:"product_id" gorm:"column:product_id"`
	Message    string `json:"message" gorm:"column:message;type:varchar(255)"`
	CreatedAt  string `json:"created_at" gorm:"column:created_at"`
	Product    Product
}

func (Komplain) TableName() string {
	return komplainTable
}
