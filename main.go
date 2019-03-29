package main

import (
	"github.com/globalsign/mgo"
	"github.com/go-chi/chi"
	"holiday-movie-bff/models"
	"holiday-movie-bff/pkg"
	"net/http"
	"time"
)

const (
	host       = "db:27017"
	database   = "hmb"
	username   = ""
	password   = ""
	collection = "squares"
)

func main() {
	r := chi.NewRouter()

	info := &mgo.DialInfo{
		Addrs:    []string{host},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	env := &handler.Env{
		DB: session,
	}

	r.Handle("/squares", handler.Handler{Env: env, H: models.GetSquares})
	defer session.Close()

	http.ListenAndServe(":3000", r)
}
