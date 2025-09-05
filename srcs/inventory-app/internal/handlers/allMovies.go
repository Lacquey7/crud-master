package handlers

import "net/http"

func (h *Handler) AllMovies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mes := "Hello World"
		w.Write([]byte(mes))
	}
}
