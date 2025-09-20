package service

import (
	"weKnow/model"
	"weKnow/repository"
)

type ReleaseServiceInterface interface {
	GetReleases() ([]model.Release, error)
}
type ReleaseService struct {
	repo *repository.ReleaseRepository
}

func NewReleaseService(repo *repository.ReleaseRepository) *ReleaseService {
	return &ReleaseService{
		repo: repo,
	}
}

func (s *ReleaseService) GetReleases() ([]model.Release, error) {
	releases, err := s.repo.GetReleases()
	if err != nil {
		return nil, err
	}
	return releases, nil
}
