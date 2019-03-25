package main

import (
	"github.com/go-chi/chi"
	"holiday-movie-bff/models/mongo"
	"net/http"
)

// App struct
type App struct {
	session *mongo.Session
}

func main() {
	r := chi.NewRouter()

	mg, err := mongo.NewSession()
	if err != nil {
		panic(err)
	}

	defer mg.Close()

	http.ListenAndServe(":3000", r)
}
