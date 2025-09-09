package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) AllMovies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("title")

		m, err := h.Service.GetMovies(title)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "Application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(m); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
