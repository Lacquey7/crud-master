package internal

import (
	"inventory-app/internal/handlers"
	"inventory-app/internal/routes"
	"inventory-app/internal/services"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	// Dependency
	DB *pgxpool.Pool

	//Services

	//Internal
	Router   *routes.Routes
	Handlers *handlers.Handler
}

func NewApp(db *pgxpool.Pool) *App {
	s := services.NewService(db)
	h := handlers.NewHandler(s)
	r := routes.NewRoutes(h)

	return &App{
		DB:       db,
		Router:   r,
		Handlers: h,
	}
}
