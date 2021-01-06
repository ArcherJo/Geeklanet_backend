package repository

import (
	"Geeklanet/datasource"
	"Geeklanet/models"
	"go.mongodb.org/mongo-driver/bson"
)

type CommentRepository struct {
	T datasource.Table
}

func convert2Comment(slice []interface{})[]models.Comment{
	var comments []models.Comment
	for _,e := range slice{
		var comment models.Comment
		bsonBytes, _ := bson.Marshal(e)
		bson.Unmarshal(bsonBytes, &comment)
		comments = append(comments,comment)
	}
	return comments
}

func (r CommentRepository) InsertOneComment(authorID,parentPostID,content string, images []string, timeStamp int64) string {
	return r.T.Insert(models.Comment{
		AuthorID:      authorID,
		ParentPostID:  parentPostID,
		Content:       content,
		ImagesBase64:  images,
		SubCommentsID: []string{},
		Like:          0,
		TimeStampUnix: timeStamp,
	})
}

func (r CommentRepository) UpdateAddLikeByID(commentID string) {
	comment := convert2Comment(r.T.Query(map[string]interface{}{
		"_id":commentID,
	}))[0]

	r.T.Update(map[string]interface{}{
		"_id":commentID,
	}, map[string]interface{}{
		"like":comment.Like+1,
	})
}


func (r CommentRepository) UpdateDecreaseLikeByID(commentID string) {
	comment := convert2Comment(r.T.Query(map[string]interface{}{
		"_id":commentID,
	}))[0]

	r.T.Update(map[string]interface{}{
		"_id":commentID,
	}, map[string]interface{}{
		"like":comment.Like-1,
	})
}

func (r CommentRepository) QueryAuthorIDByID(commentID string) string {
	res := convert2Comment(r.T.Query(map[string]interface{}{
		"_id":commentID,
	}))

	if len(res)==0{
		return ""
	}else {
		return res[0].AuthorID
	}
}

func (r CommentRepository) QuerySubCommentsIDByID(commentID string) []string {
	res := convert2Comment(r.T.Query(map[string]interface{}{
		"_id":commentID,
	}))

	if len(res)==0{
		return []string{}
	}else {
		return res[0].SubCommentsID
	}
}


func (r CommentRepository) QueryCommentByID(commentID string) models.Comment {
	res := convert2Comment(r.T.Query(map[string]interface{}{
		"_id":commentID,
	}))

	if len(res)==0{
		return models.Comment{}
	}else {
		return res[0]
	}
}

func (r CommentRepository) DeleteCommentByID(commentID string) string {
	res := convert2Comment(r.T.Delete(map[string]interface{}{
		"_id":commentID,
	}))

	if len(res)==0{
		return ""
	}else {
		return res[0].AuthorID
	}
}