package service

import (
	"weKnow/config"
	"weKnow/repository"
)

type KnownService struct {
	ArtistServiceInterface
	EventServiceInterface
	ImageServiceInterface
	ReleaseServiceInterface
}

type ServiceInterface interface {
	ArtistServiceInterface
	EventServiceInterface
	ImageServiceInterface
	ReleaseServiceInterface
}

func NewService(repo repository.KnownRepository, conf config.KnownConfig) ServiceInterface {
	return KnownService{
		ArtistServiceInterface:  NewArtistService(repo.ArtistRepositoryInterface),
		EventServiceInterface:   NewEventService(repo.EventRepositoryInterface, repo.ArtistRepositoryInterface, repo.UtilityRepositoryInterface, conf),
		ImageServiceInterface:   NewImageService(repo.ImageRepositoryInterface),
		ReleaseServiceInterface: NewReleaseService(repo.ReleaseRepositoryInterface),
	}
}
