package main

import (
	"github.com/go-chi/chi"
	"holiday-movie-bff/models"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	initMongo()
	http.ListenAndServe(":3000", r)
}
