package models

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"holiday-movie-bff/pkg"
	"net/http"
)

// Square model
type Square struct {
	ID      string `json:"_id" bson:"_id"`
	Content string `json:"content" bson:"content"`
}

// GetSquares returns all squares
func GetSquares(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	session := env.DB.Copy()
	c := session.DB("bingodb").C("squares")
	var squares []Square
	err := c.Find(bson.M{}).All(&squares)
	if err != nil {
		return handler.StatusError{Code: 500, Err: err}
	}

	json.NewEncoder(w).Encode(squares)
	return nil
}
