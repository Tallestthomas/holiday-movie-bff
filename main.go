package main

import (
	"github.com/go-chi/chi"
	"github.com/tallestthomas/holiday-movie-bff/models/mongo"
	"net/http"
)

// App struct
type App struct {
	session *mongo.Session
}

func main() {
	var a App
	var err error
	r := chi.NewRouter()
	a.session, err = mongo.NewSession()
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":3000", r)
}
