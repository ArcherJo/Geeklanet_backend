package repository

import (
	"Geeklanet/datasource"
	"Geeklanet/models"
	"go.mongodb.org/mongo-driver/bson"
)

type SubCommentRepository struct {
	T datasource.Table
}

func convert2SubComment(slice []interface{})[]models.SubComment{
	var subComments []models.SubComment
	for _,e := range slice{
		var subComment models.SubComment
		bsonBytes, _ := bson.Marshal(e)
		bson.Unmarshal(bsonBytes, &subComment)
		subComments = append(subComments,subComment)
	}
	return subComments
}

func (r SubCommentRepository) InsertOneSubComment(authorID string, parentCommentID string, content string, timeStamp int64) string {
	return r.T.Insert(models.SubComment{
		AuthorID:        authorID,
		ParentCommentID: parentCommentID,
		Content:         content,
		TimeStampUnix:   timeStamp,
	})
}


func (r SubCommentRepository) QueryAuthorIDByID(subCommentID string) string {
	res := convert2SubComment(r.T.Query(map[string]interface{}{
		"_id":subCommentID,
	}))

	if len(res)==0{
		return ""
	}else {
		return res[0].AuthorID
	}
}

func (r SubCommentRepository) DeleteSubComment(subCommentID string) {
	r.T.Delete(map[string]interface{}{
		"_id":subCommentID,
	})
}


func (r SubCommentRepository) QuerySubCommentByID(subCommentID string) models.SubComment {
	res := convert2SubComment(r.T.Query(map[string]interface{}{
		"_id":subCommentID,
	}))

	if len(res)==0{
		return models.SubComment{}
	}else {
		return res[0]
	}
}