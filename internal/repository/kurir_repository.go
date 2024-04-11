package repository

import "gorm.io/gorm"

type KurirRepository struct {
	KurirDB *gorm.DB
}

type KurirRepo interface {
}

func KurirRepoIntf(db *gorm.DB) KurirRepo {
	return &KurirRepository{
		KurirDB: db,
	}
}
