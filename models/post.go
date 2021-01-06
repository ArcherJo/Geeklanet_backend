package models

type Post struct {
	ID			string		`json:"postID" bson:"_id,omitempty"`
	Title		string		`json:"title" bson:"title"`
	AuthorID	string		`json:"authorID" bson:"authorID"`
	Tags		[]string	`json:"tags" bson:"tags"`
	CommentsID	[]string	`json:"commentsID" bson:"commentsID"`
	Type 		string		`json:"type" bson:"type"`
}