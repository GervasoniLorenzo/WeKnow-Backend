package service

import (
	"weKnow/model"
	"weKnow/repository"
)

type ReleaseServiceInterface interface {
	GetReleases() ([]model.Release, error)
	AddRelease(release model.ReleaseDto) error
	UpdateRelease(release model.ReleaseDto, id int) error
	DeleteRelease(id int) error
}
type ReleaseService struct {
	releaseRepo repository.ReleaseRepositoryInterface
	artistRepo  repository.ArtistRepositoryInterface
}

func NewReleaseService(rr repository.ReleaseRepositoryInterface, ar repository.ArtistRepositoryInterface) ReleaseServiceInterface {
	return &ReleaseService{
		releaseRepo: rr,
		artistRepo:  ar,
	}
}

func (s *ReleaseService) GetReleases() ([]model.Release, error) {
	releases, err := s.releaseRepo.GetReleases()
	if err != nil {
		return nil, err
	}
	return releases, nil
}

func (s *ReleaseService) AddRelease(releaseDto model.ReleaseDto) error {
	artists := []model.Artist{}
	for _, artistId := range releaseDto.ArtistIds {
		artist, err := s.artistRepo.GetArtistDetailsById(artistId)
		if err != nil {
			return err
		}
		artists = append(artists, artist)
	}
	release := model.Release{
		Title:       releaseDto.Title,
		ReleaseDate: releaseDto.ReleaseDate,
		Links:       releaseDto.Links,
		Artist:      artists,
	}
	return s.releaseRepo.AddRelease(release)
}

func (s *ReleaseService) UpdateRelease(releaseDto model.ReleaseDto, id int) error {
	artists := []model.Artist{}
	for _, artistId := range releaseDto.ArtistIds {
		artist, err := s.artistRepo.GetArtistDetailsById(artistId)
		if err != nil {
			return err
		}
		artists = append(artists, artist)
	}
	release := model.Release{
		Title:       releaseDto.Title,
		ReleaseDate: releaseDto.ReleaseDate,
		Links:       releaseDto.Links,
		Artist:      artists,
	}
	return s.releaseRepo.UpdateRelease(release)
}

func (s *ReleaseService) DeleteRelease(id int) error {
	return s.releaseRepo.DeleteRelease(id)
}
