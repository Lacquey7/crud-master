package internal

import (
	"billing-app/internal/handlers"
	"billing-app/internal/routes"
	"billing-app/internal/services"

	"gorm.io/gorm"
)

type App struct {
	// Dependency
	DB *gorm.DB

	//Services

	//Internal
	Router   *routes.Routes
	Handlers *handlers.Handler
}

func NewApp(db *gorm.DB) *App {
	s := services.NewService(db)
	h := handlers.NewHandler(s)
	r := routes.NewRoutes(h)

	return &App{
		DB:       db,
		Router:   r,
		Handlers: h,
	}
}
