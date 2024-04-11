package models

var tableNameLoc = "location"

type Location struct {
	Id       uint   `json:"id" gorm:"primaryKey;column:id"`
	Location string `json:"location" gorm:"column:location;type:varchar(255)"`
}

func (Location) TableName() string {
	return tableNameLoc
}
