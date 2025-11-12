package service

import (
	"fmt"
	"weKnow/model"
	"weKnow/repository"
	"weKnow/utils"
)

type ArtistServiceInterface interface {
	GetArtists() []model.ArtistBasicInfo
	AddArtist(artist model.ArtistDto) error
	UpdateArtist(artist model.ArtistDto, id int) error
	GetArtistImage(slug string) (string, string, error)
	GetArtistDetails(artistSlug string) (model.Artist, error)
	DeleteArtist(id int) error
}

type ArtistService struct {
	repo repository.ArtistRepositoryInterface
	u    utils.UtilsInterface
}

func NewArtistService(repo repository.ArtistRepositoryInterface) ArtistServiceInterface {
	return &ArtistService{
		repo: repo,
	}
}

func (s *ArtistService) GetArtists() []model.ArtistBasicInfo {
	list := []model.ArtistBasicInfo{}
	defer func() {
		n := len(list)
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if list[j].Name < list[i].Name {
					list[i], list[j] = list[j], list[i]
				}
			}
		}
	}()
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

func (s *ArtistService) AddArtist(artistDto model.ArtistDto) error {
	slug := s.u.GenerateSlug(artistDto.Name)
	exists, err := s.repo.CheckArtistSlugExists(slug)
	if err != nil {
		return err
	}
	count := 0
	for exists {
		count++
		exists, err = s.repo.CheckArtistSlugExists(fmt.Sprintf("%s-%v", slug, count))
		if err != nil {
			return err
		}
		if exists {
			slug = fmt.Sprintf("%s-%d", slug, count)
		}
	}
	artist := model.Artist{
		Name:      artistDto.Name,
		Bio:       artistDto.Bio,
		ImageUuid: &artistDto.ImageUuid,
		Slug:      slug,
	}
	return s.repo.CreateArtist(artist)
}

func (s *ArtistService) GetArtistImage(slug string) (string, string, error) {
	uuid := s.repo.GetArtistUuidBySlug(slug)
	return s.repo.GetArtistImage(uuid)
}

func (s *ArtistService) GetArtistDetails(artistSlug string) (model.Artist, error) {
	return s.repo.GetArtistDetailsBySlug(artistSlug)
}

func (s *ArtistService) UpdateArtist(artistDto model.ArtistDto, id int) error {
	artist := model.Artist{
		Id:        id,
		Name:      artistDto.Name,
		Bio:       artistDto.Bio,
		ImageUuid: &artistDto.ImageUuid,
	}
	return s.repo.UpdateArtist(artist)
}

func (s *ArtistService) DeleteArtist(id int) error {
	return s.repo.DeleteArtist(id)
}
