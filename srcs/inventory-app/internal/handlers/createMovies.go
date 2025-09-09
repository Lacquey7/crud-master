package handlers

import (
	"encoding/json"
	"inventory-app/internal/model"
	"net/http"
	"strings"
)

func (h *Handler) CreateMovies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		m := model.Movie{}

		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if strings.TrimSpace(m.Title) == "" || strings.TrimSpace(m.Description) == "" {
			http.Error(w, "Title and Description are required", http.StatusBadRequest)
			return
		}

		err = h.Service.CreateMovies(&m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}

}
