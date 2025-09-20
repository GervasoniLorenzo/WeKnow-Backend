package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"weKnow/model"
	"weKnow/service"

	"github.com/gorilla/mux"
)

type EventControllerInterface interface {
	GetNextEvent(w http.ResponseWriter, r *http.Request)
	CreateEvent(w http.ResponseWriter, r *http.Request)
	SendEventEmail(w http.ResponseWriter, r *http.Request)
}

type EventController struct {
	srv *service.EventService
}

func NewEventController(service *service.EventService) *EventController {
	return &EventController{
		srv: service,
	}
}

func (ctrl *EventController) GetNextEvent(w http.ResponseWriter, r *http.Request) {
	event, err := ctrl.srv.GetNextEvent()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false) // <--- evita la conversione
	_ = enc.Encode(event)
	// jsonData, err := json.Marshal(event)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	w.Write(buf.Bytes())
}
func (ctrl *EventController) GetPastEvents(w http.ResponseWriter, r *http.Request) {
	event, err := ctrl.srv.GetPastEvents()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	jsonData, err := json.Marshal(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
func (ctrl *EventController) GetUpcomingEvents(w http.ResponseWriter, r *http.Request) {
	event, err := ctrl.srv.GetUpcomingEvents()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	jsonData, err := json.Marshal(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
func (ctrl *EventController) CreateEvent(w http.ResponseWriter, r *http.Request) {

	event := new(model.EventDto)

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ctrl.srv.AddEvent(*event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ctrl *EventController) SendEventEmail(w http.ResponseWriter, r *http.Request) {
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

	err = ctrl.srv.SendEventEmail(eventId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ctrl *EventController) GetArtistEvents(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	slug := vars["slug"]
	if slug == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	events, err := ctrl.srv.GetArtistEvents(slug)

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
