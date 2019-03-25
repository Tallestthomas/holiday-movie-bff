package mongo

import (
	"fmt"
	"github.com/globalsign/mgo"
	"time"
)

const (
	host       = "db:27017"
	database   = "db"
	username   = ""
	password   = ""
	collection = "squares"
)

// Session DB connection session
type Session struct {
	session *mgo.Session
}

// NewSession creates a new mongodb session
func NewSession() (*Session, error) {
	info := &mgo.DialInfo{
		Addrs:    []string{host},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	fmt.Println("connected")
	if err != nil {
		panic(err)
	}
	return &Session{session}, err
}

// Copy a current Session
func (s *Session) Copy() *Session {
	return &Session{s.session.Copy()}
}

// Close a session
func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}
