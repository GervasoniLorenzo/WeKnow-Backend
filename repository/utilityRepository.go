package repository

import (
	"weKnow/adapter"
	"weKnow/db"
	"weKnow/model"
)

type UtilityRepositoryInterface interface {
	GetContacts() []model.Contact
}
type UtilityRepository struct {
	dataBase db.DatabaseInterface
	adapter  adapter.AdapterInterface
}

func NewUtilityRepository(db db.DatabaseInterface) UtilityRepositoryInterface {
	return &UtilityRepository{
		dataBase: db,
		adapter:  adapter.NewAdapter(),
	}
}

func (r *UtilityRepository) GetContacts() []model.Contact {
	return r.dataBase.GetContacts()
}
