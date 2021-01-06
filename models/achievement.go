package models

type Achievement struct {
	ID		string	`json:"achievementID" bson:"_id,omitempty"`
	Name	string	`json:"name" bson:"name"`
	Targets	[]int	`json:"targets" bson:"targets"`
}