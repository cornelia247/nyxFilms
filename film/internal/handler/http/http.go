package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/cornelia247/nyxfilms/film/internal/controller/film"
)

// Handler defines a film handler.
type Handler struct {
	ctrl *film.Controller
}

// New creates a new film HTTP handler.
func New(ctrl *film.Controller) *Handler {
	return &Handler{ctrl}
}

// GetFilmDetails handles Get /film requests.
func (h *Handler) GetFilmDetails( w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	details, err := h.ctrl.Get(req.Context(), id)
	if err != nil && errors.Is(err, film.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}else if err != nil {
		log.Printf("Repository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(details); err != nil {
		log.Printf("Reponse encode error: %v\n", err)
	}

}