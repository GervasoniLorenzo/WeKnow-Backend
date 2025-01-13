package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"weKnow/model"
	"weKnow/service"

	"github.com/gorilla/mux"
)

// KnownController rappresenta il controller principale
type KnownController struct {
	service *service.KnownService
}

// NewController crea un nuovo controller
func NewController(service *service.KnownService) *KnownController {
	return &KnownController{
		service: service,
	}
}

func (ctrl *KnownController) GetArtists(w http.ResponseWriter, r *http.Request) {

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

func (ctrl *KnownController) CreateArtist(w http.ResponseWriter, r *http.Request) {

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

func (ctrl *KnownController) GetArtistImage(w http.ResponseWriter, r *http.Request) {
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

func (ctrl *KnownController) UploadImage(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, fmt.Sprintf("Errore nel leggere il file: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()
	res, err := ctrl.service.CreateImage(handler, file)

	if err != nil {
		http.Error(w, fmt.Sprintf("Errore nel salvataggio dell'immagine: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(http.StatusOK)

	jsonData, err := json.Marshal(map[string]string{"uuid": res})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func (ctrl *KnownController) GetEventList(w http.ResponseWriter, r *http.Request) {
	events, err := ctrl.service.GetEventList()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonData, err := json.Marshal(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func (ctrl *KnownController) CreateEvent(w http.ResponseWriter, r *http.Request) {

	event := new(model.EventDto)

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ctrl.service.AddEvent(*event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ctrl *KnownController) SendEventEmail(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/event/mail/"):]

	if id == "" {
		http.Error(w, "ID event non fornito", http.StatusBadRequest)
		return
	}
	eventId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ctrl.service.SendEventEmail(eventId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ctrl *KnownController) GetArtistEvents(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	slug := vars["slug"]
	if slug == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	events, err := ctrl.service.GetArtistEvents(slug)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonData, err := json.Marshal(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
