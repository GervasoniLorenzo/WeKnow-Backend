package controller

import (
	"encoding/json"
	"net/http"
	"weKnow/service"
)

type ReleaseController struct {
	srv *service.ReleaseService
}

func NewReleaseController(service *service.ReleaseService) *ReleaseController {
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
