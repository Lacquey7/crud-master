package routes

import (
	"inventory-app/internal/handlers"

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
	prefix := r.PathPrefix("/api/movies").Subrouter()

	//Retrieve all the movies && retrieve all the movies with name in the title.
	prefix.HandleFunc("", routes.Handlers.AllMovies()).Methods("GET")

	// Create a new movie.
	prefix.HandleFunc("", routes.Handlers.CreateMovies()).Methods("POST")

	//TODO:  delete all movies in the database.
	prefix.HandleFunc("", routes.Handlers.DeleteAllMovies()).Methods("DELETE")

	//TODO: retrieve a single movie by id.
	prefix.HandleFunc("/{id}", routes.Handlers.GetMovie()).Methods("GET")

	//TODO: update a single movie by id.
	prefix.HandleFunc("/{id}", routes.Handlers.UpdateMovie()).Methods("PUT")

	//TODO: delete a single movie by id.
	prefix.HandleFunc("/{id}", routes.Handlers.DeleteMovie()).Methods("DELETE")

}
