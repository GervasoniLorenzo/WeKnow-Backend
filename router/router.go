package router

import (
	"weKnow/controller"
	"weKnow/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter(ctrl controller.KnownController) *mux.Router {
	r := mux.NewRouter()

	// jwtSecret := []byte(os.Getenv("APP_JWT_HS256_SECRET"))
	
	// ARTISTS
	r.HandleFunc("/artist/list", ctrl.GetArtists).Methods("GET")
	r.HandleFunc("/artist/image/{slug}", ctrl.GetArtistImage).Methods("GET")
	r.HandleFunc("/artist/{slug}/events", ctrl.GetArtistEvents).Methods("GET")
	// r.HandleFunc("/artist/image", ctrl.UploadImage).Methods("POST")
	r.HandleFunc("/artist", ctrl.CreateArtist).Methods("POST")
	r.HandleFunc("/artist/{slug}", ctrl.GetArtistDetails).Methods("GET")

	// EVENTS
	r.HandleFunc("/event/next", ctrl.GetNextEvent).Methods("GET")
	r.HandleFunc("/event/upcoming", ctrl.GetUpcomingEvents).Methods("GET")
	r.HandleFunc("/event/past", ctrl.GetPastEvents).Methods("GET")
	r.HandleFunc("/event/mail/{id}", ctrl.SendEventEmail).Methods("POST")
	r.HandleFunc("/event/image/{slug}", ctrl.GetEventImage).Methods("GET")

	// RELEASES
	r.HandleFunc("/release/list", ctrl.GetReleases).Methods("GET")
	r.HandleFunc("/release/image", ctrl.GetReleaseImage).Methods("GET")

	// ADMIN
	admin := r.PathPrefix("/admin").Subrouter()
	// admin.Use(middleware.AppJWT(jwtSecret))
	admin.Use(middleware.AdminOnly())
	admin.HandleFunc("/event/list", ctrl.AdminGetEventList).Methods("GET")
	admin.HandleFunc("/event", ctrl.AdminCreateEvent).Methods("POST")
	admin.HandleFunc("/event/{id}", ctrl.AdminUpdateEvent).Methods("PUT")
	admin.HandleFunc("/event/{id}", ctrl.AdminDeleteEvent).Methods("DELETE")
	admin.HandleFunc("/event/image", ctrl.UploadEventImage).Methods("POST")

	admin.HandleFunc("/artist/list", ctrl.GetArtists).Methods("GET")
	// admin.HandleFunc("/artist/{slug}", ctrl.UpdateArtist).Methods("PUT")
	// admin.HandleFunc("/artist/{slug}", ctrl.DeleteArtist).Methods("DELETE")
	// admin.HandleFunc("/release", ctrl.CreateRelease).Methods("POST")
	// admin.HandleFunc("/release/{id}", ctrl.UpdateRelease).Methods("PUT")
	// admin.HandleFunc("/release/{id}", ctrl.DeleteRelease).Methods("DELETE")
	return r
}
