package repository

import (
	"Geeklanet/datasource"
	"Geeklanet/models"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {
	T datasource.Table
}

func convert2User(slice []interface{})[]models.User{
	var users []models.User
	for _,e := range slice{
		var user models.User
		bsonBytes, _ := bson.Marshal(e)
		bson.Unmarshal(bsonBytes, &user)
		users = append(users,user)
	}
	return users
}

func (r *UserRepository) QueryUserIDByName(userName string) string {
	res := convert2User(r.T.Query(map[string]interface{}{
		"name":userName,
	}))

	if len(res)==0{
		return ""
	}else {
		return res[0].ID
	}
}

func (r *UserRepository) QueryPasswordByID(userID string) string {
	res := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))

	if len(res)==0{
		return ""
	}else {
		return res[0].Password
	}
}

func (r *UserRepository) QueryPasswordByName(userName string) string {
	res := convert2User(r.T.Query(map[string]interface{}{
		"name":userName,
	}))
	if len(res)==0{
		return ""
	}else {
		return res[0].Password
	}
}

func (r *UserRepository) InsertOneUser(name string, password string, email string, gender bool) string{
	return r.T.Insert(models.User{
		Name: name,
		Password: password,
		Email: email,
		Status: "",
		Gender: gender,
		Avatar: "",
		Achievement: []int{},
		Age: 0,
		Admin:false,
		Level: 0,
		LikedCount: 0,

		Posts:		[]string{},
		Comments:	[]string{},
		SubComments:[]string{},
		Liked:		[]string{},
		Disliked:	[]string{},
		Favorites:	[]string{},
		Follows:	[]string{},
		Followers:	[]string{},
		Notices:	[]string{},

	})
}

func (r *UserRepository) UpdateAddLikedCommentByID(userID string, commentID string) {
	user := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))[0]

	user.Liked = append(user.Liked, commentID)

	r.T.Update(map[string]interface{}{
			"_id":userID,
		},
		map[string]interface{}{
			"liked":user.Liked,
		},
	)
}

func (r *UserRepository) UpdateAddDislikedCommentByID(userID string, commentID string) {
	user := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))[0]

	user.Disliked = append(user.Disliked, commentID)

	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"disliked":user.Disliked,
		},
	)
}

func (r *UserRepository) UpdateIncreaseLikeCountByID(userID string) {
	user := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))[0]


	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"likedCount":user.LikedCount+1,
		},
	)
}

func (r *UserRepository) UpdateDecreaseLikeCountByID(userID string) {
	user := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))[0]


	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"likedCount":user.LikedCount-1,
		},
	)
}

func (r *UserRepository) UpdateAddPostByID(userID string, postID string) {
	user := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))[0]

	user.Posts = append(user.Posts, postID)

	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"posts":user.Posts,
		},
	)
}

func (r *UserRepository) UpdateAddCommentByID(userID string, commentID string) {
	user := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))[0]

	user.Comments = append(user.Comments, commentID)

	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"comments":user.Comments,
		},
	)
}

func (r *UserRepository) UpdateAddSubCommentByID(userID string, subCommentID string) {
	user := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))[0]

	user.SubComments = append(user.SubComments, subCommentID)

	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"subComments":user.SubComments,
		},
	)
}

func (r *UserRepository) UpdateRemovePostByID(userID string, postID string) {
	user := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))[0]

	for i,id:= range user.Posts{
		if id == postID{
			user.Posts=append(user.Posts[:i],user.Posts[i+1:]...)
			break
		}
	}

	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"posts":user.Posts,
		},
	)
}

func (r *UserRepository) UpdateRemoveCommentByID(userID string, commentID string) {
	user := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))[0]

	for i,id:= range user.Comments{
		if id == commentID{
			user.Comments=append(user.Comments[:i],user.Comments[i+1:]...)
			break
		}
	}

	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"comments":user.Comments,
		},
	)
}

func (r *UserRepository) UpdateRemoveSubCommentByID(userID string, subCommentID string) {
	user := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))[0]

	for i,id:= range user.SubComments{
		if id == subCommentID{
			user.SubComments=append(user.SubComments[:i],user.SubComments[i+1:]...)
			break
		}
	}

	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"subComments":user.SubComments,
		},
	)
}

func (r *UserRepository) UpdateRemoveLikedCommentByID(userID string, commentID string) {
	user := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))[0]

	for i,id:= range user.Liked{
		if id == commentID{
			user.Liked=append(user.Liked[:i],user.Liked[i+1:]...)
			break
		}
	}

	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"liked":user.Liked,
		},
	)
}

func (r *UserRepository) UpdateRemoveDisLikedCommentByID(userID string, commentID string) {
	user := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))[0]

	for i,id:= range user.Liked{
		if id == commentID{
			user.Disliked=append(user.Disliked[:i],user.Disliked[i+1:]...)
			break
		}
	}

	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"disliked":user.Disliked,
		},
	)
}

func (r *UserRepository) QueryFollowsByID(userID string) []string {
	res := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))

	if len(res)==0{
		return []string{}
	}else {
		return res[0].Follows
	}

}

func (r *UserRepository) QueryFollowersByID(userID string) []string {
	res := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))
	if len(res)==0{
		return []string{}
	}else {
		return res[0].Followers
	}
}

func (r *UserRepository) QueryUserByID(userID string) models.User {
	res := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))
	if len(res)==0{
		return models.User{}
	}else {
		return res[0]
	}
}

func (r *UserRepository) UpdateAddFavoriteCommentByID(userID string, commentID string) {
	user := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))[0]

	user.Favorites = append(user.Favorites, commentID)

	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"favorites":user.Favorites,
		},
	)
}

func (r *UserRepository) UpdateRemoveFavoriteCommentByID(userID string, commentID string) {
	user := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))[0]

	for i,id:= range user.Favorites{
		if id == commentID{
			user.Favorites=append(user.Favorites[:i],user.Favorites[i+1:]...)
			break
		}
	}

	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"favorites":user.Favorites,
		},
	)
}

func (r *UserRepository) UpdateModifyNameByID(userID string, userName string) {
	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"name":userName,
		},
	)
}

func (r *UserRepository) UpdateModifyStatusByID(userID string, status string) {
	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"status":status,
		},
	)
}

func (r *UserRepository) UpdateModifyAvatarByID(userID string, avatar string) {
	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"avatar":avatar,
		},
	)
}

func (r *UserRepository) UpdateAddFollowByID(userID string, userID2 string) {
	user := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))[0]

	user.Follows = append(user.Follows, userID2)

	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"follows":user.Follows,
		},
	)
}

func (r *UserRepository) UpdateRemoveFollowByID(userID string, userID2 string) {
	user := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))[0]

	for i,id:= range user.Followers{
		if id == userID2{
			user.Followers=append(user.Followers[:i],user.Favorites[i+1:]...)
			break
		}
	}

	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"followers":user.Followers,
		},
	)
}

func (r *UserRepository) UpdateAddNoticeByID(userID string, noticeID string) {
	user := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))[0]

	user.Notices = append(user.Follows, noticeID)

	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"notices":user.Notices,
		},
	)
}


func (r *UserRepository) UpdateRemoveNoticeByID(userID string, noticeID string) {
	user := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))[0]

	for i,id:= range user.Followers{
		if id == noticeID{
			user.Notices=append(user.Notices[:i],user.Notices[i+1:]...)
			break
		}
	}

	r.T.Update(map[string]interface{}{
		"_id":userID,
	},
		map[string]interface{}{
			"notices":user.Notices,
		},
	)
}

func (r *UserRepository) QueryPostsIDByID(userID string) []string{
	res := convert2User(r.T.Query(map[string]interface{}{
		"_id":userID,
	}))

	if len(res)==0{
		return []string{}
	}else {
		return res[0].Posts
	}
}
