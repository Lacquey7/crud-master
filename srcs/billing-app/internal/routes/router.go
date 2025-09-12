package routes

import (
	"billing-app/internal/handlers"

	"github.com/gorilla/mux"
)

type Routes struct {
	Handlers *handlers.Handler
}

func NewRoutes(h *handlers.Handler) *Routes {
	return &Routes{
		Handlers: h,
	}
}

func (routes *Routes) Router(r *mux.Router) {
	prefix := r.PathPrefix("/api/billing").Subrouter()

	prefix.HandleFunc("/{id}", routes.Handlers.GetBillingById()).Methods("GET")
}
