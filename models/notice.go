package models

type NoticeType int

const (
	NoticeTypeFollow = iota
	NoticeTypeLike
	NoticeTypeDislike
	NoticeTypeComment
	NoticeTypeSubComment
	NoticeTypeCallInComment
	NoticeTypeCallInSubComment
)

type Notice struct {
	ID					string		`json:"noticeID" bson:"_id,omitempty"`
	Type				NoticeType	`json:"type,int" bson:"type,int"`
	RecipientID			string		`json:"recipientID" bson:"recipientID"`
	SponsorID			string		`json:"sponsorID" bson:"sponsorID"`
	PostID				string		`json:"postID" bson:"postID"`
	CommentID			string		`json:"commentID" bson:"commentID"`
	SubCommentID		string		`json:"subCommentID" bson:"subCommentID"`
	PostTitle			string		`json:"postTitle" bson:"postTitle"`
	CommentContent		string		`json:"commentContent" bson:"commentContent"`
	SubCommentContent	string		`json:"subCommentContent" bson:"subCommentContent"`
	HasRead				bool		`json:"hasRead" bson:"hasRead"`
	TimeStampUnix		int64		`json:"timeStampUnix,string" bson:"timeStampUnix,string"`
}
