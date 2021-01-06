package models

type Tag struct {
	ID		string	`json:"tagID" bson:"_id,omitempty"`
	Name	string	`json:"name" bson:"name"`
}
