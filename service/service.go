package service

import (
	"mime/multipart"
	"weKnow/repository"
)

type KnownService struct {
	*ArtistService
	*EventService
	*ImageService
	*ReleaseService
}

func NewService(repo *repository.KnownRepository) *KnownService {
	return &KnownService{
		ArtistService:  NewArtistService(repo.ArtistRepository),
		EventService:   NewEventService(repo.EventRepository, repo.ArtistRepository, repo.UtilityRepository),
		ImageService:   NewImageService(),
		ReleaseService: NewReleaseService(repo.ReleaseRepository),
	}
}

type ImageServiceInterface interface {
	CreateImage(handler *multipart.FileHeader, file multipart.File) (string, error)
}
