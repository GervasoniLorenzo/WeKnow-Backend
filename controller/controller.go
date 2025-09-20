package controller

import (
	"weKnow/service"
)

type KnownController struct {
	*EventController
	*ArtistController
	*ImageController
	*ReleaseController
}

func NewController(service *service.KnownService) *KnownController {
	return &KnownController{
		EventController:   NewEventController(service.EventService),
		ArtistController:  NewArtistController(service.ArtistService),
		ImageController:   NewImageController(service.ImageService),
		ReleaseController: NewReleaseController(service.ReleaseService),
	}
}
