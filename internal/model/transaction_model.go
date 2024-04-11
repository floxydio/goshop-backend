package models

var transactionTable = "transaction_history"

type TransactionHistory struct {
	TransactionID  uint   `json:"id" gorm:"primaryKey;column:transaction_id"`
	TrxCode        string `json:"trx_code" gorm:"column:trx_code;type:varchar(255)"`
	UserID         uint   `json:"user_id" gorm:"column:user_id"`
	ProductID      uint   `json:"product_id" gorm:"column:product_id"`
	BuyDate        string `json:"buy_date" gorm:"column:buy_date"`
	StatusBuy      uint   `json:"status_buy" gorm:"column:status_buy;type:int(11);default:0"`
	Quantity       uint   `json:"quantity" gorm:"column:quantity;type:int(11);default:0"`
	TotalPrice     uint   `json:"total_price" gorm:"column:total_price;type:int(30);default:0"`
	KurirServiceID uint   `json:"kurir_service_id" gorm:"column:kurir_service_id"`
	KurirUsersID   uint   `json:"kurir_users" gorm:"column:kurir_users"`
	StasusSending  uint   `json:"status_sending" gorm:"column:status_sending;type:int(11);default:0"`
	CreatedAt      string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt      string `json:"updated_at" gorm:"column:updated_at; default:NULL"`
	User           User
	Product        Product
	KurirService   KurirService
	KurirUsers     UserKurir
}

func (TransactionHistory) TableName() string {
	return transactionTable
}
