package controller

import (
	"Geeklanet/models"
	"Geeklanet/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"math/rand"
	"time"
)
type PostController struct {
	S service.Service
}

func (c PostController)getPostsBaseInfo(postsID []string) mvc.Result{
	postsInfo, usersID := c.S.Post.GetPostBaseInfoSortedByTimeByIDs(postsID)
	var usersInfo []models.User
	for _,userID := range usersID{
		userInfo := c.S.User.GetUserBaseInfo(userID)
		usersInfo = append(usersInfo, userInfo)
	}

	var posts []map[string]interface{}
	for i := range postsInfo{
		posts = append(posts, map[string]interface{}{
			"Author":usersInfo[i],
			"Post":postsInfo[i],
		})
	}
	return mvc.Response{
		Object: posts,
	}
}

func (c PostController)GetPersonalBy(userID string) mvc.Result{
	postsID := c.S.User.GetPersonalPostByID(userID)

	return c.getPostsBaseInfo(postsID)
}


func (c PostController)GetFollow(context iris.Context) mvc.Result{
	userID := sessions.Get(context).GetString("userID")
	followsID := c.S.User.GetFollowsID(userID)
	var postsID []string
	for _,userID := range followsID{
		post := c.S.User.GetPersonalPostByID(userID)
		postsID = append(postsID, post...)
	}
	return c.getPostsBaseInfo(postsID)

}


func (c PostController)GetPersonal(context iris.Context) mvc.Result{
	session := sessions.Get(context)
	userID := session.GetString("userID")
	postsID := c.S.User.GetPersonalPostByID(userID)
	return mvc.Response{
		Object: map[string]interface{}{
			"Data":c.getPostsBaseInfo(postsID),
		},
	}
}


func (c PostController)GetRecommend(context iris.Context) mvc.Result{
	iris.New().Logger().Info(context.Path())
	session := sessions.Get(context)
	userID := session.GetString("userID")
	recommender := c.S.Recommend.GetRecommendWeight(userID)
	return mvc.Response{
		Object: map[string]interface{}{
			"Data":c.S.Post.GetRecommendPostBaseInfoByRecommender(recommender),
		},
	}

}


func (c PostController)GetPopular() mvc.Result{
	return mvc.Response{
		Object: map[string]interface{}{
			"Data":c.S.Post.GetPopularPost(func(post models.Post) float64 {
				return float64(rand.Intn(100)) / 100.0
			}),
		},
	}
}





func (c PostController)PostPost(context iris.Context){
	receive := struct {
		PostTitle				string		`json:"title"`
		PostTags				[]string	`json:"tags"`
		PostType 				string		`json:"type"`

		CommentContent			string		`json:"content"`
		CommentImagesBase64		[]string	`json:"imagesBase64"`
	}{}

	if err := context.ReadJSON(&receive);err!=nil{
		context.StopWithError(iris.StatusBadRequest, err)
		return
	}

	userID := sessions.Get(context).GetString("userID")
	postID,commentID := c.S.Post.CreatePost(userID, receive.PostTitle, receive.PostTags,
										  receive.PostType, receive.CommentContent,
										  receive.CommentImagesBase64, time.Now().Unix())

	usersCalledID := c.S.User.CreatePost(userID, postID, receive.CommentContent)

	for _,userCalledID := range usersCalledID{
		noticeID := c.S.Notice.CreateCallInComment(userCalledID,userID,postID,commentID,receive.PostTitle,receive.CommentContent,time.Now().Unix())
		c.S.User.GetNotice(userCalledID,noticeID)
	}

}

func (c PostController)PostComment(context iris.Context){
	receive := struct {
		Content			string		`json:"content"`
		ImagesBase64	[]string	`json:"imagesBase64"`
		ParentPostID	string		`json:"parentPostID"`
	}{}

	if err := context.ReadJSON(&receive);err!=nil{
		context.StopWithError(iris.StatusBadRequest, err)
		return
	}

	userID := sessions.Get(context).GetString("userID")
	commentID := c.S.Post.CreateComment(userID, receive.ParentPostID, receive.Content, receive.ImagesBase64, time.Now().Unix())
	userCommentedID := c.S.Post.GetPostAuthorByID(receive.ParentPostID)
	postInfo := c.S.Post.GetPostByID(receive.ParentPostID)

	usersCalledID := c.S.User.CreateComment(userID,commentID, receive.Content)

	noticeCommentedID := c.S.Notice.CreateComment(userCommentedID,userID,receive.ParentPostID,commentID,postInfo.Title,receive.Content,time.Now().Unix())

	for _,userCalledID := range usersCalledID{
		noticeID := c.S.Notice.CreateCallInComment(userCalledID,userID,receive.ParentPostID,commentID,postInfo.Title,receive.Content,time.Now().Unix())
		c.S.User.GetNotice(userCalledID,noticeID)
	}

	c.S.User.GetNotice(userCommentedID, noticeCommentedID)
}


func (c PostController)PostSubcomment(context iris.Context){
	receive := struct {
		Content			string		`json:"content"`
		ParentCommentID	string		`json:"parentCommentID"`
	}{}

	if err := context.ReadJSON(&receive);err!=nil{
		context.StopWithError(iris.StatusBadRequest, err)
		return
	}

	iris.New().Logger().Info(context.Path())
	userID := sessions.Get(context).GetString("userID")

	subCommentID := c.S.Post.CreateSubComment(userID, receive.ParentCommentID,receive.Content, time.Now().Unix())
	userSubCommentedID := c.S.Post.GetCommentAuthor(receive.ParentCommentID)
	commentInfo := c.S.Post.GetCommentByID(receive.ParentCommentID)
	postInfo := c.S.Post.GetPostByID(commentInfo.ParentPostID)

	usersCalledID := c.S.User.CreateSubComment(userID,subCommentID,receive.Content)

	noticeSubCommentedID := c.S.Notice.CreateSubComment(userSubCommentedID,userID,postInfo.ID,commentInfo.ID,subCommentID,postInfo.Title,commentInfo.Content,receive.Content,time.Now().Unix())
	for _,userCalledID := range usersCalledID{
		noticeID := c.S.Notice.CreateCallInSubComment(userCalledID,userID,postInfo.ID,commentInfo.ID,subCommentID,postInfo.Title,commentInfo.Content,receive.Content,time.Now().Unix())
		c.S.User.GetNotice(userCalledID,noticeID)
	}
	c.S.User.GetNotice(userSubCommentedID, noticeSubCommentedID)
}

func (c PostController) DeletePost(postID string,context iris.Context)  {
	userID := sessions.Get(context).GetString("userID")
	if userID != c.S.Post.GetPostAuthorByID(postID){
		context.StatusCode(iris.StatusMethodNotAllowed)
		return
	}
	c.deletePost(postID)

}

func (c PostController) deletePost(postID string)  {
	for _,commentID := range c.S.Post.GetCommentsID(postID){
		c.deleteComment(commentID)
	}
	userID :=c.S.Post.DeletePost(postID)
	c.S.User.DeletePost(userID,postID)
	noticesID,recipientsID := c.S.Notice.DeletePost(postID)
	for i,_ := range noticesID{
		c.S.User.DeleteNotice(recipientsID[i],noticesID[i])
	}
}



func (c PostController) DeleteComment(commentID string,context iris.Context)  {
	userID := sessions.Get(context).GetString("userID")
	if userID != c.S.Post.GetCommentAuthor(commentID){
		context.StatusCode(iris.StatusMethodNotAllowed)
		return
	}
	c.deleteComment(commentID)

}

func (c *PostController) deleteComment(commentID string){
	for _,subCommentID := range c.S.Post.GetSubCommentsID(commentID){
		c.deleteSubcomment(subCommentID)
	}
	userID:=c.S.Post.DeleteComment(commentID)
	c.S.User.DeleteComment(userID,commentID)
	noticesID,recipientsID := c.S.Notice.DeleteComment(commentID)
	for i,_ := range noticesID{
		c.S.User.DeleteNotice(recipientsID[i],noticesID[i])
	}
}


func (c *PostController) DeleteSubcomment(subcommentID string,context iris.Context)  {
	userID := sessions.Get(context).GetString("userID")
	if userID != c.S.Post.GetSubCommentAuthorByID(subcommentID){
		context.StatusCode(iris.StatusMethodNotAllowed)
		return
	}
	c.deleteSubcomment(subcommentID)
}

func (c *PostController) deleteSubcomment(subcommentID string){
	userID := c.S.Post.DeleteSubComment(subcommentID)
	c.S.User.DeleteSubComment(userID,subcommentID)
	noticesID,recipientsID := c.S.Notice.DeleteSubComment(subcommentID)
	for i,_ := range noticesID{
		c.S.User.DeleteNotice(recipientsID[i],noticesID[i])
	}
}


func (c *PostController) DeleteLike(commentID string,context iris.Context)  {
	sponsorID := sessions.Get(context).GetString("userID")
	if c.S.Notice.HasLiked(sponsorID,commentID){
		context.StatusCode(iris.StatusMethodNotAllowed)
		return
	}
	c.S.User.DeleteLikeComment(sponsorID,commentID)
	c.S.Post.DeleteLike(commentID)
	noticeID,recipientID := c.S.Notice.DeleteLike(sponsorID,commentID)
	c.S.User.DeleteBeenLiked(recipientID)
	c.S.User.DeleteNotice(recipientID,noticeID)

}


func (c *PostController) DeleteDislike(commentID string,context iris.Context)  {
	sponsorID := sessions.Get(context).GetString("userID")
	if c.S.Notice.HasDisliked(sponsorID,commentID){
		context.StatusCode(iris.StatusMethodNotAllowed)
		return
	}
	c.S.User.DeleteDislikeComment(sponsorID,commentID)
	c.S.Post.DeleteDislike(commentID)
	noticeID,recipientID := c.S.Notice.DeleteDislike(sponsorID,commentID)
	c.S.User.DeleteBeenDisliked(recipientID)
	c.S.User.DeleteNotice(recipientID,noticeID)

}

func (c* PostController)GetDetailBy(postID string)mvc.Result{
	return mvc.Response{
		Object: c.S.Post.GetPostDetailByID(postID),
	}
}

func (c* PostController)PostLikeBy(commentID string, context iris.Context){
	sponsorID := sessions.Get(context).GetString("userID")

	recipientID := c.S.Post.LikeCommentByID(commentID)
	commentInfo := c.S.Post.GetCommentByID(commentID)
	postInfo := c.S.Post.GetPostByID(commentInfo.ParentPostID)

	c.S.User.LikeComment(sponsorID, commentID)
	c.S.User.BeenLiked(recipientID)

	noticeID := c.S.Notice.LikeNotice(recipientID, sponsorID, postInfo.ID,commentID,postInfo.Title,commentInfo.Content,time.Now().Unix())

	c.S.User.GetNotice(recipientID,noticeID)
}


func (c* PostController)PostDislikeBy(commentID string, context iris.Context){
	sponsorID := sessions.Get(context).GetString("userID")

	recipientID := c.S.Post.DislikeCommentByID(commentID)
	commentInfo := c.S.Post.GetCommentByID(commentID)
	postInfo := c.S.Post.GetPostByID(commentInfo.ParentPostID)

	c.S.User.DislikeComment(sponsorID, commentID)
	c.S.User.BeenDisliked(recipientID)

	noticeID := c.S.Notice.DislikeNotice(recipientID, sponsorID, postInfo.ID,commentID,postInfo.Title,commentInfo.Content,time.Now().Unix())

	c.S.User.GetNotice(recipientID,noticeID)
}

func (c PostController) PostFavoriteBy(commentID string, context iris.Context)  {
	userID := sessions.Get(context).GetString("userID")
	c.S.User.FavoriteComment(userID,commentID)
}


func (c PostController) DeleteFavoriteBy(commentID string, context iris.Context)  {
	userID := sessions.Get(context).GetString("userID")
	c.S.User.DeleteFavoriteComment(userID,commentID)
}

