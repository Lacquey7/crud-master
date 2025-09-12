package handlers

import "api-gateway/internal/services"

type Handler struct {
	Service *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{
		Service: services,
	}
}
