package handlers

import "net/http"

func (h *Handler) DeleteAllMovies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := h.Service.DeleteMovies()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
