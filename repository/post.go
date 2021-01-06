package repository

import (
	"Geeklanet/datasource"
	"Geeklanet/models"
	"go.mongodb.org/mongo-driver/bson"
)

type PostRepository struct {
	T datasource.Table
}

func convert2Post(slice []interface{})[]models.Post{
	var posts []models.Post
	for _,e := range slice{
		var post models.Post
		bsonBytes, _ := bson.Marshal(e)
		bson.Unmarshal(bsonBytes, &post)
		posts = append(posts,post)
	}
	return posts
}

func (r PostRepository) InsertOnePost(authorID string, title string, tags []string, postType string) string {
	return r.T.Insert(models.Post{
		Title:      title,
		AuthorID:   authorID,
		Tags:       tags,
		CommentsID: []string{},
		Type:       postType,
	})
}

func (r PostRepository) QueryPostByID(postID string) models.Post {
	res := convert2Post(r.T.Query(map[string]interface{}{
		"postID":postID,
	}))

	if len(res)==0{
		return models.Post{}
	}else {
		return res[0]
	}
}

func (r PostRepository) QueryAuthorIDByID(postID string) string {
	res := convert2Post(r.T.Query(map[string]interface{}{
		"postID":postID,
	}))

	if len(res)==0{
		return ""
	}else {
		return res[0].AuthorID
	}
}

func (r PostRepository) QueryCommentsIDByID(postID string) []string {
	res := convert2Post(r.T.Query(map[string]interface{}{
		"postID":postID,
	}))

	if len(res)==0{
		return []string{}
	}else {
		return res[0].CommentsID
	}

}

func (r PostRepository) DeletePostByID(postID string) models.Post {
	res := convert2Post(r.T.Delete(map[string]interface{}{
		"postID":postID,
	}))

	if len(res)==0{
		return models.Post{}
	}else {
		return res[0]
	}
}

func (r PostRepository) QueryPost() []models.Post {
	return convert2Post(r.T.Query(map[string]interface{}{}))
}

