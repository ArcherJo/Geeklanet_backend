package models

type User struct {
	ID			string		`json:"userID" bson:"_id,omitempty"`
	Name		string		`json:"name" bson:"name"`
	Password	string		`json:"password" bson:"password"`
	Email		string		`json:"email" bson:"email"`
	Gender		bool		`json:"gender" bson:"gender"`
	Status		string		`json:"status" bson:"status"`
	Avatar		string		`json:"avatar" bson:"avatar"`
	Achievement	[]int		`json:"achievement" bson:"achievement"`
	Age			int			`json:"age" bson:"age"`
	Admin		bool		`json:"admin" bson:"admin"`
	Level		int			`json:"level" bson:"level"`
	LikedCount	int			`json:"likedCount" bson:"likedCount"`

	Posts		[]string	`json:"posts" bson:"posts"`
	Comments	[]string	`json:"comments" bson:"comments"`
	SubComments	[]string	`json:"subComments" bson:"subComments"`
	Liked		[]string	`json:"liked" bson:"liked"`
	Disliked	[]string	`json:"disliked" bson:"disliked"`
	Favorites	[]string	`json:"favorites" bson:"favorites"`
	Follows		[]string	`json:"follows" bson:"follows"`
	Followers	[]string	`json:"followers" bson:"followers"`
	Notices		[]string	`json:"notices" bson:"notices"`
}