package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Server initializes chi routes
func Server() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	return router
}
