package main

import (
	"github.com/globalsign/mgo"
	"github.com/go-chi/chi"
	"holiday-movie-bff/models"
	"holiday-movie-bff/models/mongo"
	"net/http"
)

// App struct
type App struct {
	session *mgo.Session
}

func main() {
	r := chi.NewRouter()

	var app App
	var err error
	app.session, err = mongo.NewSession()
	if err != nil {
		panic(err)
	}

	r.HandleFunc("/squares", models.GetSquares(session))
	defer session.Close()

	http.ListenAndServe(":3000", r)
}
