package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"weKnow/model"
	"weKnow/service"

	"github.com/gorilla/mux"
)

type ArtistController struct {
	service service.ArtistServiceInterface
}

type ArtistControllerInterface interface {
	GetArtists(w http.ResponseWriter, r *http.Request)
	CreateArtist(w http.ResponseWriter, r *http.Request)
	GetArtistDetails(w http.ResponseWriter, r *http.Request)
	UpdateArtist(w http.ResponseWriter, r *http.Request)
	DeleteArtist(w http.ResponseWriter, r *http.Request)
}

func NewArtistController(service service.ArtistServiceInterface) ArtistControllerInterface {
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

	artist := new(model.ArtistDto)
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

func (ctrl *ArtistController) UpdateArtist(w http.ResponseWriter, r *http.Request) {
	artistDto := new(model.ArtistDto)
	err := json.NewDecoder(r.Body).Decode(&artistDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok || idStr == "" {
		http.Error(w, "missing artist id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "invalid artist id", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}
	err = ctrl.service.UpdateArtist(*artistDto, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ctrl *ArtistController) DeleteArtist(w http.ResponseWriter, r *http.Request) {
vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok || idStr == "" {
		http.Error(w, "missing artist id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "invalid artist id", http.StatusBadRequest)
		return
	}
	err = ctrl.service.DeleteArtist(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
