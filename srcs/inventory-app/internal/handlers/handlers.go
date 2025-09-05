package handlers

import "inventory-app/internal/services"

type Handler struct {
	S services.Service
}

func NewHandler(services services.Service) *Handler {
	return &Handler{
		S: services,
	}
}
