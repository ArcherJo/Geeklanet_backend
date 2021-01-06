package models

type SubComment struct {
	ID 				string	`json:"subCommentID" bson:"_id,omitempty"`
	AuthorID 		string	`json:"authorID" bson:"authorID"`
	ParentCommentID	string	`json:"parentCommentID" bson:"parentCommentID"`
	Content			string	`json:"content" bson:"content"`
	TimeStampUnix	int64	`json:"timeStampUnix,string" bson:"timeStampUnix,string"`
}