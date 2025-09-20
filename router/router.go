package router

import (
	"weKnow/controller"

	"github.com/gorilla/mux"
)

func SetupRouter(ctrl *controller.KnownController) *mux.Router {
	router := mux.NewRouter()

	// ARTISTS
	router.HandleFunc("/artist/list", ctrl.GetArtists).Methods("GET")
	router.HandleFunc("/artist/image/{slug}", ctrl.GetArtistImage).Methods("GET")
	router.HandleFunc("/artist/{slug}/events", ctrl.GetArtistEvents).Methods("GET")
	router.HandleFunc("/artist/image", ctrl.UploadImage).Methods("POST")
	router.HandleFunc("/artist", ctrl.CreateArtist).Methods("POST")
	router.HandleFunc("/artist/{slug}", ctrl.GetArtistDetails).Methods("GET")

	// EVENTS
	router.HandleFunc("/event/next", ctrl.GetNextEvent).Methods("GET")
	router.HandleFunc("/event/upcoming", ctrl.GetUpcomingEvents).Methods("GET")
	router.HandleFunc("/event/past", ctrl.GetPastEvents).Methods("GET")
	router.HandleFunc("/event", ctrl.CreateEvent).Methods("POST")
	router.HandleFunc("/event/mail/{id}", ctrl.SendEventEmail).Methods("POST")
	router.HandleFunc("/event/image", ctrl.GetEventImage).Methods("GET")
	
	// RELEASES
	router.HandleFunc("/release/list", ctrl.GetReleases).Methods("GET")
	router.HandleFunc("/release/image", ctrl.GetReleaseImage).Methods("GET")

	return router
}
