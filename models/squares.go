package models

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/chi"
	"holiday-movie-bff/pkg"
	"net/http"
)

// Square model
type Square struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Content string        `json:"content" bson:"content"`
}

// GetSquares returns all squares
func GetSquares(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	session := env.DB.Copy()
	defer session.Close()
	c := session.DB("bingodb").C("squares")
	var squares []Square
	err := c.Find(bson.M{}).All(&squares)
	if err != nil {
		return handler.StatusError{Code: 500, Err: err}
	}

	json.NewEncoder(w).Encode(squares)
	return nil
}

// GetSquare returns a square by id
func GetSquare(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	session := env.DB.Copy()
	defer session.Close()
	c := session.DB("bingodb").C("squares")

	param := chi.URLParam(r, "id")
	var square Square
	err := c.FindId(bson.ObjectIdHex(param)).One(&square)
	if err != nil {
		return handler.StatusError{Code: 500, Err: err}
	}

	json.NewEncoder(w).Encode(square)
	return nil
}

// AddSquare creates a new square
func AddSquare(env *handler.Env, w http.ResponseWriter, r *http.Request) error {

	session := env.DB.Copy()
	defer session.Close()
	c := session.DB("bingodb").C("squares")

	var square Square
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&square)

	if err != nil {
		return handler.StatusError{Code: 500, Err: err}
	}

	err = c.Insert(&square)
	if err != nil {
		return handler.StatusError{Code: 500, Err: err}
	}

	return nil
}
