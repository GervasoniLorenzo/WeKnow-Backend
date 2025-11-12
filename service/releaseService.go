package service

import (
	"fmt"
	"strings"
	"weKnow/model"
	"weKnow/repository"
	"weKnow/utils"
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
	u           utils.UtilsInterface
}

func NewReleaseService(rr repository.ReleaseRepositoryInterface, ar repository.ArtistRepositoryInterface) ReleaseServiceInterface {
	return &ReleaseService{
		releaseRepo: rr,
		artistRepo:  ar,
		u:           utils.NewUtils(),
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
	slug := s.u.GenerateSlug(releaseDto.Title)
	exists, err := s.releaseRepo.CheckReleaseSlugExists(slug)
	if err != nil {
		return err
	}
	count := 0
	for exists {
		count++
		exists, err = s.releaseRepo.CheckReleaseSlugExists(fmt.Sprintf("%s-%v", slug, count))
		if err != nil {
			return err
		}
		if exists {
			slug = fmt.Sprintf("%s-%d", slug, count)
		}
	}
	for _, artistId := range releaseDto.ArtistIds {
		artist, err := s.artistRepo.GetArtistDetailsById(artistId)
		if err != nil {
			return err
		}
		artists = append(artists, artist)
	}
	release := model.Release{
		Title:   releaseDto.Title,
		Date:    releaseDto.Date,
		Links:   releaseDto.Links,
		Artists: artists,
	}
	return s.releaseRepo.AddRelease(release)
}

func (s *ReleaseService) UpdateRelease(releaseDto model.ReleaseDto, id int) error {

	if len(releaseDto.ArtistIds) == 0 {
		return fmt.Errorf("artistsIds is required and cannot be empty")
	}
	artists := make([]model.Artist, 0, len(releaseDto.ArtistIds))
	for _, aid := range releaseDto.ArtistIds {
		artists = append(artists, model.Artist{Id: aid})
	}

	links := make([]model.ReleaseLink, 0, len(releaseDto.Links))
	for _, l := range releaseDto.Links {
		if l.ID == 0 {
			l.ReleaseID = id
		}
		links = append(links, l)
	}

	release := model.Release{
		ID:      id,
		Title:   strings.TrimSpace(releaseDto.Title),
		Date:    releaseDto.Date,
		Label:   strings.TrimSpace(releaseDto.Label),
		Artists: artists,
		Links:   links,
	}

	return s.releaseRepo.UpdateRelease(release)
}

func (s *ReleaseService) DeleteRelease(id int) error {
	return s.releaseRepo.DeleteRelease(id)
}
