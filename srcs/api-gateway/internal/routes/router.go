package routes

import (
	"api-gateway/internal/handlers"
	"api-gateway/pkg"

	"github.com/gorilla/mux"
)

type Routes struct {
	Handlers *handlers.Handler
	RabbitMq *pkg.RabbitMQ
}

func NewRoutes(h *handlers.Handler, mq *pkg.RabbitMQ) *Routes {
	return &Routes{
		Handlers: h,
		RabbitMq: mq,
	}
}

func (routes *Routes) Router(r *mux.Router) {
	billing := r.PathPrefix("/billing").Subrouter()
	inventory := r.PathPrefix("/inventory").Subrouter()

	billing.HandleFunc("", routes.Handlers.BillingApi()).Methods("GET")
	inventory.HandleFunc("", routes.Handlers.InventoryApi()).Methods("GET")
	
}
