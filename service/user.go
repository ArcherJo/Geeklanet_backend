package service

import (
	"Geeklanet/models"
	"Geeklanet/repository"
	"strings"
)

type userService struct {
	r repository.UserRepository
}

func (s *userService)CheckUserName(userName string) bool{
	userID := s.r.QueryUserIDByName(userName)

	return userID != ""
}


func (s *userService) GetUserIDByName(name string) string {
	return s.r.QueryUserIDByName(name)
}

func (s *userService)GetUserPasswordByID(userID string) string{
	return s.r.QueryPasswordByID(userID)
}
func (s *userService)GetUserPasswordByName(userName string) string{
	return s.r.QueryPasswordByName(userName)
}
func (s *userService)CreateNewAccount(userName, userPassword, userEmail string, gender bool) string{
	if s.CheckUserName(userName){
		return ""
	}
	return s.r.InsertOneUser(userName, userPassword, userEmail ,gender)
}

func (s *userService) LikeComment(userID, commentID string) {
	s.r.UpdateAddLikedCommentByID(userID, commentID)
}

func (s *userService) DislikeComment(userID, commentID string) {
	s.r.UpdateAddDislikedCommentByID(userID, commentID)
}

func (s *userService) BeenLiked(userID string) {
	s.r.UpdateIncreaseLikeCountByID(userID)
}

func (s *userService) BeenDisliked(userID string) {
	s.r.UpdateDecreaseLikeCountByID(userID)
}

func (s *userService) CreatePost(userID, postID, content string) []string{
	s.r.UpdateAddPostByID(userID, postID)
	return s.getCall(content)
}


func (s *userService) CreateComment(userID, commentID, content string) []string{
	s.r.UpdateAddCommentByID(userID, commentID)
	return s.getCall(content)
}

func (s *userService) CreateSubComment(userID, subCommentID, content string) []string{
	s.r.UpdateAddSubCommentByID(userID, subCommentID)
	return s.getCall(content)
}

func (s *userService) DeletePost(userID, postID string) {
	s.r.UpdateRemovePostByID(userID, postID)
}

func (s *userService) DeleteComment(userID, commentID string) {
	s.r.UpdateRemoveCommentByID(userID, commentID)
}

func (s *userService) DeleteSubComment(userID, subCommentID string) {
	s.r.UpdateRemoveSubCommentByID(userID, subCommentID)
}

func (s *userService) DeleteLikeComment(userID, commentID string) {
	s.r.UpdateRemoveLikedCommentByID(userID,commentID)
}

func (s *userService) DeleteDislikeComment(userID, commentID string) {
	s.r.UpdateRemoveDisLikedCommentByID(userID,commentID)
}

func (s *userService) DeleteBeenLiked(userID string) {
	s.r.UpdateDecreaseLikeCountByID(userID)
}

func (s *userService) DeleteBeenDisliked(userID string) {
	s.r.UpdateIncreaseLikeCountByID(userID)
}

func (s *userService) GetFollowsID(userID string) []string {
	return s.r.QueryFollowsByID(userID)
}

func (s *userService) GetFollowersID(userID string) []string {
	return s.r.QueryFollowersByID(userID)
}

func (s *userService) GetUserBaseInfo(userID string) models.User {
	return s.r.QueryUserByID(userID)
}

func (s *userService) FavoriteComment(userID, commentID string) {
	s.r.UpdateAddFavoriteCommentByID(userID, commentID)
}


func (s *userService) DeleteFavoriteComment(userID, commentID string) {
	s.r.UpdateRemoveFavoriteCommentByID(userID, commentID)

}

func (s *userService) ModifyName(userID, userName string) bool {
	if s.CheckUserName(userName){
		return false
	}
	s.r.UpdateModifyNameByID(userID, userName)
	return true
}

func (s *userService) ModifyStatus(userID, status string) {
	s.r.UpdateModifyStatusByID(userID, status)
}


func (s *userService) ModifyAvatar(userID, avatar string) {
	s.r.UpdateModifyAvatarByID(userID, avatar)
}

func (s *userService) Follow(recipientID,sponsorID string) {
	s.r.UpdateAddFollowByID(recipientID, sponsorID)
}

func (s *userService) DeleteFollow(recipientID,sponsorID string) {
	s.r.UpdateRemoveFollowByID(recipientID, sponsorID)
}

func (s *userService) GetNotice(userID, noticeID string) {
	s.r.UpdateAddNoticeByID(userID, noticeID)
}

func (s *userService) GetPersonalPostByID(userID string) []string {
	return s.r.QueryPostsIDByID(userID)
}

func (s *userService) getCall(content string) []string {
	var usersName, usersID []string

	fragments := strings.Split(content," ")
	for _,str := range fragments{
		if strings.HasPrefix(str,"@"){
			usersName = append(usersName,strings.TrimPrefix(str,"@"))
		}
	}
	for _,userName := range usersName{
		usersID = append(usersID,s.r.QueryUserIDByName(userName))
	}
	return usersID
}

func (s *userService) DeleteNotice(userID, noticeID string) {
	s.r.UpdateRemoveNoticeByID(userID, noticeID)
}



