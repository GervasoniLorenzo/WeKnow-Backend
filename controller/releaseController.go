package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"weKnow/model"
	"weKnow/service"

	"github.com/gorilla/mux"
)

type ReleaseController struct {
	srv service.ReleaseServiceInterface
}

type ReleaseControllerInterface interface {
	GetReleases(w http.ResponseWriter, r *http.Request)
}

func NewReleaseController(service service.ReleaseServiceInterface) ReleaseControllerInterface {
	return &ReleaseController{
		srv: service,
	}
}
func (ctrl *ReleaseController) GetReleases(w http.ResponseWriter, r *http.Request) {
	releases, err := ctrl.srv.GetReleases()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(http.StatusOK)

	jsonData, err := json.Marshal(releases)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func (ctrl *ReleaseController) AddRelease(w http.ResponseWriter, r *http.Request) {
	release := new(model.ReleaseDto)
	err := json.NewDecoder(r.Body).Decode(&release)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ctrl.srv.AddRelease(*release)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ctrl *ReleaseController) UpdateRelease(w http.ResponseWriter, r *http.Request) {
	release := new(model.ReleaseDto)
	err := json.NewDecoder(r.Body).Decode(&release)
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
		http.Error(w, "Invalid release ID", http.StatusBadRequest)
		return
	}

	err = ctrl.srv.UpdateRelease(*release, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ctrl *ReleaseController) DeleteRelease(w http.ResponseWriter, r *http.Request) {
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

	err = ctrl.srv.DeleteRelease(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
