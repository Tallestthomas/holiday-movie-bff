package main

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		fmt.Print(err)
	}

	dbs, dbErr := session.DatabaseNames()
	if dbErr != nil {
		fmt.Print(err)
	}
	fmt.Print(dbs)
}
