package controller

import (
	"weKnow/service"
)

type KnownController struct {
	EventControllerInterface
	ArtistControllerInterface
	ImageControllerInterface
	ReleaseControllerInterface
}

type ControllerInterface interface {
	EventControllerInterface
	ArtistControllerInterface
	ImageControllerInterface
	ReleaseControllerInterface
}

func NewController(service service.KnownService) ControllerInterface {
	return KnownController{
		EventControllerInterface:   NewEventController(service.EventServiceInterface),
		ArtistControllerInterface:  NewArtistController(service.ArtistServiceInterface),
		ImageControllerInterface:   NewImageController(service.ImageServiceInterface),
		ReleaseControllerInterface: NewReleaseController(service.ReleaseServiceInterface),
	}
}
