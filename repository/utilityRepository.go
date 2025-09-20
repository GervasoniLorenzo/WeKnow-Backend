package repository

import (
	"weKnow/adapter"
	"weKnow/db"
	"weKnow/model"
)

type UtilityRepositoryInterface interface {
	GetContacts()
}
type UtilityRepository struct {
	dataBase *db.KnownDatabase
	adapter  *adapter.KnownAdapter
}

func NewUtilityRepository(db *db.KnownDatabase) *UtilityRepository {
	return &UtilityRepository{
		dataBase: db,
		adapter:  adapter.NewAdapter(),
	}
}

func (r *UtilityRepository) GetContacts() []model.Contact {
	return r.dataBase.GetContacts()
}
