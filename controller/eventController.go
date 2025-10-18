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
	SendEventEmail(w http.ResponseWriter, r *http.Request)
	AdminGetEventList(w http.ResponseWriter, r *http.Request)
	AdminDeleteEvent(w http.ResponseWriter, r *http.Request)
	AdminUpdateEvent(w http.ResponseWriter, r *http.Request)
	AdminCreateEvent(w http.ResponseWriter, r *http.Request)
	GetPastEvents(w http.ResponseWriter, r *http.Request)
	GetUpcomingEvents(w http.ResponseWriter, r *http.Request)
	GetArtistEvents(w http.ResponseWriter, r *http.Request)
}

type EventController struct {
	srv service.EventServiceInterface
}

func NewEventController(service service.EventServiceInterface) EventControllerInterface {
	return &EventController{
		srv: service,
	}
}

func (ctrl *EventController) GetNextEvent(w http.ResponseWriter, r *http.Request) {
	view := r.URL.Query().Get("view")
	event, err := ctrl.srv.GetNextEvent(view)

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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
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

func (ctrl *EventController) AdminGetEventList(w http.ResponseWriter, r *http.Request) {
	events, err := ctrl.srv.AdminGetEventList()

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

func (ctrl *EventController) AdminDeleteEvent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/admin/event/"):]

	if id == "" {
		http.Error(w, "ID event non fornito", http.StatusBadRequest)
		return
	}
	eventId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ctrl.srv.AdminDeleteEvent(eventId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ctrl *EventController) AdminUpdateEvent(w http.ResponseWriter, r *http.Request) {

	event := new(model.UpdateEventDto)

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ctrl.srv.AdminUpdateEvent(*event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ctrl *EventController) AdminCreateEvent(w http.ResponseWriter, r *http.Request) {

	event := new(model.EventDto)

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ctrl.srv.AdminCreateEvent(*event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
