package internal

import (
	"api-gateway/internal/handlers"
	"api-gateway/internal/routes"
	"api-gateway/internal/services"
	"api-gateway/pkg"
)

type App struct {
	//Internal
	Router   *routes.Routes
	Handlers *handlers.Handler
	RabbitMQ *pkg.RabbitMQ
}

func NewApp(mq *pkg.RabbitMQ) *App {
	s := services.NewService()
	h := handlers.NewHandler(s)
	r := routes.NewRoutes(h, mq)

	return &App{
		Router:   r,
		Handlers: h,
		RabbitMQ: mq,
	}
}
