package repository

import (
	"weKnow/db"
	"weKnow/model"
)

type ReleaseRepositoryInterface interface {
	GetReleases() ([]model.Release, error)
}
type ReleaseRepository struct {
	dataBase *db.KnownDatabase
}

func NewReleaseRepository() *ReleaseRepository {
	return &ReleaseRepository{}
}
func (r *ReleaseRepository) GetReleases() ([]model.Release, error) {
	res, err := r.dataBase.GetReleases()
	if err != nil {
		return nil, err
	}
	return res, nil
}
