package models

type Product struct {
	ProductID uint   `json:"id" gorm:"primaryKey;column:product_id"`
	Name      string `json:"name" gorm:"column:name;type:varchar(255)"`
	Category  string `json:"category" gorm:"column:category;type:varchar(50)"`
	Price     uint   `json:"pricing" gorm:"column:price"`
	Images    string `json:"product_image" gorm:"column:product_image;type:varchar(255)"`
	Quantity  uint   `json:"quantity" gorm:"column:quantity"`
}

func (Product) TableName() string {
	return "product"
}
