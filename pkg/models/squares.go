package models

// Square model
type Square struct {
	ID      string `json:"id" bson:"id"`
	Content string `json:"content" bson:"content"`
}
