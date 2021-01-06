package models

type Recommend struct {
	ID		string	`json:"ID" bson:"_id,omitempty"`
	UserID	string	`json:"userID" bson:"userID"`
	Weights	[]int	`json:"weights" bson:"weights"`
}