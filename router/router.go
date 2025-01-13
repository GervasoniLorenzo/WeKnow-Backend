package router

import (
	"weKnow/controller"

	"github.com/gorilla/mux"
)

func SetupRouter(ctrl *controller.KnownController) *mux.Router {
	router := mux.NewRouter()

	// Definisci le rotte e i relativi handler
	router.HandleFunc("/artist/list", ctrl.GetArtists).Methods("GET")
	router.HandleFunc("/artist/image/{slug}", ctrl.GetArtistImage).Methods("GET")
	router.HandleFunc("/artist/{slug}/events", ctrl.GetArtistEvents).Methods("GET")
	router.HandleFunc("/artist/image", ctrl.UploadImage).Methods("POST")
	router.HandleFunc("/artist", ctrl.CreateArtist).Methods("POST")
	router.HandleFunc("/event/list", ctrl.GetEventList).Methods("GET")
	router.HandleFunc("/event", ctrl.CreateEvent).Methods("POST")
	router.HandleFunc("/event/mail/{id}", ctrl.SendEventEmail).Methods("POST")
	return router
}
