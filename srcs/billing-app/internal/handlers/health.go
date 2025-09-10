package handlers

import "net/http"

func (h *Handler) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "OK"
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(msg))
	}
}
