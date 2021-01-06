package models

type Comment struct {
	ID 				string		`json:"commentID" bson:"_id,omitempty"`
	AuthorID 		string		`json:"authorID" bson:"authorID"`
	ParentPostID	string		`json:"parentPostID" bson:"parentPostID"`
	Content			string		`json:"content" bson:"content"`
	ImagesBase64	[]string	`json:"imagesBase64" bson:"imagesBase64"`
	SubCommentsID	[]string	`json:"subCommentsID" bson:"subCommentsID"`
	Like			int			`json:"like" bson:"like"`
	TimeStampUnix	int64		`json:"timeStampUnix,string" bson:"timeStampUnix,string"`
}