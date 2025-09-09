package handlers

import (
	"encoding/json"
	"inventory-app/internal/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) UpdateMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m1 model.Movie

		if err := json.NewDecoder(r.Body).Decode(&m1); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		idStr := mux.Vars(r)["id"]

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		m, err := h.Service.GetMoviesById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		m1.ID = id
		err = h.Service.UpdateMovies(&m1, m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	}
}
