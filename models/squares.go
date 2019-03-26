package models

import (
	"encoding/json"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"net/http"
)

// Square model
type Square struct {
	ID      string `json:"_id" bson:"_id"`
	Content string `json:"content" bson:"content"`
}

// GetSquares returns all squares
func GetSquares(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	session := s.Copy()

	c := session.DB("bingodb").C("squares")
	var squares []Square
	err := c.Find(bson.M{}).All(&squares)
	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(squares)
	}
}
