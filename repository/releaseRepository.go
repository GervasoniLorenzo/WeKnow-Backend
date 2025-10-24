package repository

import (
	"weKnow/db"
	"weKnow/model"
)

type ReleaseRepositoryInterface interface {
	GetReleases() ([]model.Release, error)
	AddRelease(model.Release) error
	UpdateRelease(release model.Release) error
	DeleteRelease(id int) error
}
type ReleaseRepository struct {
	dataBase db.DatabaseInterface
}

func NewReleaseRepository(db db.DatabaseInterface) ReleaseRepositoryInterface {
	return &ReleaseRepository{
		dataBase: db,
	}
}
func (r *ReleaseRepository) GetReleases() ([]model.Release, error) {
	res, err := r.dataBase.GetReleases()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *ReleaseRepository) AddRelease(release model.Release) error {
	return r.dataBase.CreateRelease(release)
}

func (r *ReleaseRepository) UpdateRelease(release model.Release) error {
	return r.dataBase.UpdateRelease(release)
}
func (r *ReleaseRepository) DeleteRelease(id int) error {
	return r.dataBase.DeleteRelease(id)
}
