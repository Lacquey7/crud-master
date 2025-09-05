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
	prefix := r.PathPrefix("/api/movies")

	//TODO : retrieve all the movies && retrieve all the movies with name in the title.
	prefix.HandlerFunc(routes.Handlers.AllMovies()).Methods("GET")

	//TODO: create a new product entry.
	prefix.HandlerFunc(routes.Handlers.CreateMovies()).Methods("POST")

	//TODO:  delete all movies in the database.
	prefix.HandlerFunc(routes.Handlers.DeleteAllMovies()).Methods("DELETE")

	//TODO: retrieve a single movie by id.
	prefix.Path("/{id}").HandlerFunc(routes.Handlers.GetMovie()).Methods("GET")

	//TODO: update a single movie by id.
	prefix.Path("/{id}").HandlerFunc(routes.Handlers.UpdateMovie()).Methods("PUT")

	//TODO: delete a single movie by id.
	prefix.Path("/{id}").HandlerFunc(routes.Handlers.DeleteMovie()).Methods("DELETE")

}
