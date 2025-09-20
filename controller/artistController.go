package controller

import (
	"encoding/json"
	"net/http"
	"weKnow/model"
	"weKnow/service"
)

type ArtistController struct {
	service *service.ArtistService
}

type ArtistControllerInterface interface {
	GetArtists(w http.ResponseWriter, r *http.Request)
	AddArtist(w http.ResponseWriter, r *http.Request)
	GetArtistImage(w http.ResponseWriter, r *http.Request)
}

func NewArtistController(service *service.ArtistService) *ArtistController {
	return &ArtistController{
		service: service,
	}
}

func (ctrl *ArtistController) GetArtists(w http.ResponseWriter, r *http.Request) {

	artists := ctrl.service.GetArtists()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonData, err := json.Marshal(artists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func (ctrl *ArtistController) CreateArtist(w http.ResponseWriter, r *http.Request) {

	artist := new(model.Artist)
	err := json.NewDecoder(r.Body).Decode(&artist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ctrl.service.AddArtist(*artist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ctrl *ArtistController) GetArtistImage(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[len("/artist/image/"):]
	if slug == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	img, mimetype, err := ctrl.service.GetArtistImage(slug)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", mimetype)
	http.ServeFile(w, r, img)
}

func (ctrl *ArtistController) GetArtistDetails(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[len("/artist/"):]
	if slug == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	details, err := ctrl.service.GetArtistDetails(slug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonData, err := json.Marshal(details)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
