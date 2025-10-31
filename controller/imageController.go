package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weKnow/service"

	"github.com/gorilla/mux"
)

type ImageController struct {
	srv service.ImageServiceInterface
}
type ImageControllerInterface interface {
	UploadEventImage(w http.ResponseWriter, r *http.Request)
	GetEventImage(w http.ResponseWriter, r *http.Request)
	// GetArtistImage(w http.ResponseWriter, r *http.Request)
	GetReleaseImage(w http.ResponseWriter, r *http.Request)
	UploadReleaseImage(w http.ResponseWriter, r *http.Request)
}

func NewImageController(service service.ImageServiceInterface) ImageControllerInterface {
	return &ImageController{
		srv: service,
	}
}

func (ctrl *ImageController) UploadEventImage(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, fmt.Sprintf("Errore nel leggere il file: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()
	res, err := ctrl.srv.CreateImage(handler, file, "event")

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

func (ctrl *ImageController) GetEventImage(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[len("/event/image/"):]
	if slug == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	img, mimetype, err := ctrl.srv.GetEventImage(slug)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", mimetype)
	http.ServeFile(w, r, img)
}

// func (ctrl *ImageController) GetArtistImage(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Query().Get("id")
// 	artistType := r.URL.Query().Get("type")
// 	if id == "" || artistType == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	img, mimetype, err := ctrl.srv.GetArtistsImage(id, artistType)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}
// 	w.Header().Set("Content-Type", mimetype)
// 	http.ServeFile(w, r, img)
// }

func (ctrl *ImageController) GetReleaseImage(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	if slug == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	img, mimetype, err := ctrl.srv.GetReleaseImage(slug)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", mimetype)
	http.ServeFile(w, r, img)
}

func (ctrl *ImageController) UploadReleaseImage(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, fmt.Sprintf("Errore nel leggere il file: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()
	res, err := ctrl.srv.CreateImage(handler, file, "release")

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
