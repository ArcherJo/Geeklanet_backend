package service

import (
	"Geeklanet/models"
	"Geeklanet/repository"
)

type postService struct {
	postR repository.PostRepository
	commentR repository.CommentRepository
	subCommentR repository.SubCommentRepository
}



func (s *postService) GetPostDetailByID(postID string) interface{} {
	type comment struct {
		Comment 	models.Comment		`json:"commentHead"`
		SubComments	[]models.SubComment	`json:"subComments"`
	}
	type post struct {
		Post	models.Post		`json:"postHead"`
		Comments	[]comment		`json:"comments"`
	}

	var postDetail post
	postDetail.Post = s.postR.QueryPostByID(postID)
	for _,commentID := range postDetail.Post.CommentsID{
		var commentDetail comment
		commentDetail.Comment = s.commentR.QueryCommentByID(commentID)
		for _,subCommentID := range commentDetail.Comment.SubCommentsID{
			commentDetail.SubComments = append(commentDetail.SubComments, s.subCommentR.QuerySubCommentByID(subCommentID))
		}
		postDetail.Comments = append(postDetail.Comments, commentDetail)
	}

	return postDetail
}

func (s *postService) LikeCommentByID(commentID string) string{
	s.commentR.UpdateAddLikeByID(commentID)
	return s.commentR.QueryAuthorIDByID(commentID)
}

func (s *postService) DislikeCommentByID(commentID string) string {
	s.commentR.UpdateDecreaseLikeByID(commentID)
	return s.commentR.QueryAuthorIDByID(commentID)
}

func (s *postService) DeleteSubComment(subCommentID string) string {
	userID := s.subCommentR.QueryAuthorIDByID(subCommentID)
	s.subCommentR.DeleteSubComment(subCommentID)
	return userID
}

func (s *postService) GetSubCommentAuthorByID(subCommentID string) string {
	return s.subCommentR.QueryAuthorIDByID(subCommentID)
}

func (s *postService) GetPostAuthorByID(postID string) string {
	return s.postR.QueryAuthorIDByID(postID)
}

func (s *postService) GetCommentsID(postID string) []string{
	return s.postR.QueryCommentsIDByID(postID)
}

func (s *postService) DeletePost(postID string) string {
	return s.postR.DeletePostByID(postID).AuthorID
}

func (s *postService) GetCommentAuthor(commentID string) string {
	return s.commentR.QueryAuthorIDByID(commentID)
}

func (s *postService) GetSubCommentsID(commentID string) []string{
	return s.commentR.QuerySubCommentsIDByID(commentID)
}

func (s *postService) DeleteComment(commentID string) string {
	return s.commentR.DeleteCommentByID(commentID)
}

func (s *postService) DeleteLike(commentID string) {
	s.commentR.UpdateDecreaseLikeByID(commentID)
}

func (s *postService) DeleteDislike(commentID string) {
	s.commentR.UpdateAddLikeByID(commentID)
}

func (s *postService) GetPostBaseInfoSortedByTimeByIDs(postsID []string) ([]interface{}, []string) {
	type postBase struct {
		ID 				string 		`json:"postID"`
		Tags			[]string	`json:"tags"`
		Title			string 		`json:"title"`
		Type			string 		`json:"type"`
		Like			int 		`json:"like"`
		CommentCount	int 		`json:"commentCount"`
		Content			string 		`json:"content"`
		TimeStampUnix	int64		`json:"timeStampUnix,string"`
	}
	var posts []interface{}
	var usersID []string
	for _,postID := range postsID{
		postInfo := s.postR.QueryPostByID(postID)
		commentInfo := s.commentR.QueryCommentByID(postInfo.CommentsID[0])
		post := postBase{
			ID: postID,
			Tags: postInfo.Tags,
			Title: postInfo.Title,
			Type: postInfo.Type,
			Like: commentInfo.Like,
			CommentCount: len(postInfo.CommentsID),
			Content: commentInfo.Content,
			TimeStampUnix: commentInfo.TimeStampUnix,
		}
		posts = append(posts, post)
		usersID = append(usersID, postInfo.AuthorID)
	}
	return posts,usersID
}

func (s *postService) GetRecommendPostBaseInfoByRecommender(recommender func(post models.Post) float64) []models.Post {
	var posts []models.Post
	posts = s.postR.QueryPost()
	var recommendedPosts []models.Post
	for _,post := range posts{
		if recommender(post) >=0.5{
			recommendedPosts = append(recommendedPosts,post)
		}
	}

	return recommendedPosts
}

func (s *postService) CreatePost(authorID string, title string, tags []string, postType string, content string, images []string, timeStamp int64) (string,string) {
	postID := s.postR.InsertOnePost(authorID,title,tags,postType)

	return postID,s.CreateComment(authorID,postID,content,images,timeStamp)
}

func (s *postService) CreateComment(authorID string, parentPostID string, content string, images []string, timeStamp int64) string {
	return s.commentR.InsertOneComment(authorID,parentPostID,content,images,timeStamp)
}

func (s *postService) CreateSubComment(authorID string, parentCommentID string, content string, timeStamp int64) string {
	return s.subCommentR.InsertOneSubComment(authorID,parentCommentID,content,timeStamp)
}

func (s *postService) GetPopularPost(popular func(post models.Post) float64) []models.Post {
	var posts []models.Post
	posts = s.postR.QueryPost()
	var popularPosts []models.Post
	for _,post := range posts{
		if popular(post) >=0.5{
			popularPosts = append(popularPosts,post)
		}
	}

	return popularPosts
}

func (s *postService) GetPostByID(postID string) models.Post {
	return s.postR.QueryPostByID(postID)
}

func (s *postService) GetCommentByID(commentID string) models.Comment {
	return s.commentR.QueryCommentByID(commentID)
}

