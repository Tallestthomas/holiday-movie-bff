package main

import (
	"github.com/globalsign/mgo"
	"github.com/go-chi/chi"
	"net/http"
	"time"
)

const (
	host       = "db:27017"
	database   = "db"
	username   = ""
	password   = ""
	collection = "squares"
)

func initMongo() (session *mgo.Session) {
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
	return session
}

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}
