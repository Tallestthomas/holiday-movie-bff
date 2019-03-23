package main

import (
	"fmt"
	"github.com/globalsign/mgo"
	"time"
)

const (
	host       = "localhost:27017"
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

	fmt.Println("initializing")
	session, err := mgo.DialWithInfo(info)
	fmt.Println("connected")
	if err != nil {
		panic(err)
	}
	return session
}

func main() {
	mongodb := initMongo()
	fmt.Println("Mongodb Initialized")
	dbs, err := mongodb.DatabaseNames()
	if err != nil {
		panic(err)
	}
	fmt.Println(dbs)
}
