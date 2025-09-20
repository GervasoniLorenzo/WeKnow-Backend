package service

import (
	"weKnow/model"
	"weKnow/repository"
)

type ArtistServiceInterface interface {
	GetArtists() []model.Artist
	AddArtist(artist model.Artist) error
	GetArtistImage(slug string) (string, string, error)
}

type ArtistService struct {
	repo *repository.ArtistRepository
}

func NewArtistService(repo *repository.ArtistRepository) *ArtistService {
	return &ArtistService{
		repo: repo,
	}
}

func (s *ArtistService) GetArtists() []model.ArtistBasicInfo {
	list := []model.ArtistBasicInfo{}
	for _, artist := range s.repo.GetArtists() {
		artist := model.ArtistBasicInfo{
			Id:   artist.Id,
			Name: artist.Name,
			Slug: artist.Slug,
		}
		list = append(list, artist)
	}
	return list
}

func (s *ArtistService) AddArtist(artist model.Artist) error {
	return s.repo.CreateArtist(artist)
}

func (s *ArtistService) GetArtistImage(slug string) (string, string, error) {
	uuid := s.repo.GetArtistUuidBySlug(slug)
	return s.repo.GetArtistImage(uuid)
}

func (s *ArtistService) GetArtistDetails(artistSlug string) (model.Artist, error) {
	return s.repo.GetArtistDetailsBySlug(artistSlug)
}