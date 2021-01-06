package service

import (
	"Geeklanet/models"
	"Geeklanet/repository"
)

type noticeService struct {
	r repository.NoticeRepository
}

func (s *noticeService) HasLiked(sponsorID,commentID string) bool {
	return s.r.QueryNoticeIDBySponsorIDandCommentIDandType(models.NoticeTypeLike,sponsorID,commentID) !=""
}

func (s *noticeService) HasDisliked(sponsorID,commentID string) bool {
	return s.r.QueryNoticeIDBySponsorIDandCommentIDandType(models.NoticeTypeDislike,sponsorID,commentID) !=""
}


func (s *noticeService) DeleteFollow(sponsorID,recipientID string) string{
	noticeID := s.r.QueryNoticeIDBySponsorIDandRecipientIDandType(models.NoticeTypeFollow,sponsorID,recipientID)
	s.r.DeleteNoticeByID(noticeID)
	return noticeID
}

func (s *noticeService) DeletePost(postID string) ([]string,[]string){
	notices := s.r.DeleteNoticeByPostID(postID)
	var noticesID, usersID []string
	for _,notice := range notices{
		noticesID = append(noticesID,notice.ID)
		usersID = append(usersID,notice.RecipientID)
	}
	return noticesID,usersID
}

func (s *noticeService) DeleteComment(commentID string) ([]string,[]string){
	notices := s.r.DeleteNoticeByCommentID(commentID)
	var noticesID, usersID []string
	for _,notice := range notices{
		noticesID = append(noticesID,notice.ID)
		usersID = append(usersID,notice.RecipientID)
	}
	return noticesID,usersID
}

func (s *noticeService) DeleteSubComment(subCommentID string)([]string,[]string) {
	notices := s.r.DeleteNoticeBySubCommentID(subCommentID)
	var noticesID, usersID []string
	for _,notice := range notices{
		noticesID = append(noticesID,notice.ID)
		usersID = append(usersID,notice.RecipientID)
	}
	return noticesID,usersID
}


func (s *noticeService) DeleteLike(sponsorID,commentID string) (string ,string){
	noticeID := s.r.QueryNoticeIDBySponsorIDandCommentIDandType(models.NoticeTypeLike,sponsorID,commentID)
	recipientID := s.r.QueryRecipientIDByNoticeID(noticeID)
	s.r.DeleteNoticeByID(noticeID)
	return noticeID,recipientID
}

func (s *noticeService) DeleteDislike(sponsorID,commentID string) (string ,string){
	noticeID := s.r.QueryNoticeIDBySponsorIDandCommentIDandType(models.NoticeTypeDislike,sponsorID,commentID)
	recipientID := s.r.QueryRecipientIDByNoticeID(noticeID)
	s.r.DeleteNoticeByID(noticeID)
	return noticeID,recipientID
	
}


func (s *noticeService) GetUnreadNoticeByRecipientID(recipientID string) []models.Notice {
	return s.r.QueryNoticeByRecipientIDandUnread(recipientID)
}

func (s *noticeService) GetFollowNoticeByRecipientID(recipientID string) []models.Notice {
	return s.r.QueryNoticeIDByRecipientIDandType(models.NoticeTypeFollow,recipientID)

}

func (s *noticeService) merge(notice1 []models.Notice,notice2 []models.Notice)[]models.Notice{
	var res []models.Notice
	for len(notice1)==0 || len(notice2)==0{
		if notice1[0].TimeStampUnix > notice2[0].TimeStampUnix{
			res = append(res,notice1[0])
			notice1=notice1[1:]
		}else {
			res = append(res,notice2[0])
			notice2=notice2[1:]
		}
	}
	if len(notice1)==0{
		res = append(res,notice2...)
	}else {
		res = append(res,notice1...)
	}
	return res
}

func (s *noticeService) GetLikeNoticeByRecipientID(recipientID string) []models.Notice {
	likeNotice := s.r.QueryNoticeIDByRecipientIDandType(models.NoticeTypeLike,recipientID)
	dislikeNotice := s.r.QueryNoticeIDByRecipientIDandType(models.NoticeTypeDislike,recipientID)
	return s.merge(likeNotice,dislikeNotice)
}

func (s *noticeService) GetCallNoticeByRecipientID(recipientID string) []models.Notice {
	callInCommentNotice := s.r.QueryNoticeIDByRecipientIDandType(models.NoticeTypeCallInComment,recipientID)
	callInSubCommentNotice := s.r.QueryNoticeIDByRecipientIDandType(models.NoticeTypeCallInSubComment,recipientID)
	return s.merge(callInCommentNotice,callInSubCommentNotice)
}

func (s *noticeService) GetReplyNoticeByRecipientID(recipientID string) interface{} {
	commentNotice := s.r.QueryNoticeIDByRecipientIDandType(models.NoticeTypeComment,recipientID)
	subCommentNotice := s.r.QueryNoticeIDByRecipientIDandType(models.NoticeTypeSubComment,recipientID)
	return s.merge(commentNotice,subCommentNotice)
}



func (s *noticeService) Follow(sponsorID,recipientID string,TimeStamp int64) string{
	return s.r.InsertNotice(models.NoticeTypeFollow,recipientID,sponsorID,"","","","","","",TimeStamp)
}

func (s *noticeService) CreateComment(recipientID string, sponsorID string, postID string, commentID string, title string, commentContent string, TimeStamp int64) string {
	return s.r.InsertNotice(models.NoticeTypeComment,recipientID,sponsorID,postID,commentID,"",title,commentContent,"",TimeStamp)
}

func (s *noticeService) CreateSubComment(recipientID string, sponsorID string, postID string, commentID string, subCommentID string, title string, commentContent string, subCommentContent string, TimeStamp int64) string {
	return s.r.InsertNotice(models.NoticeTypeSubComment,recipientID,sponsorID,postID,commentID,subCommentID,title,commentContent,subCommentContent,TimeStamp)
}

func (s *noticeService) CreateCallInComment(recipientID string, sponsorID string, postID string, commentID string, title string, commentContent string, TimeStamp int64) string {
	return s.r.InsertNotice(models.NoticeTypeCallInComment,recipientID,sponsorID,postID,commentID,"",title,commentContent,"",TimeStamp)
}


func (s *noticeService) CreateCallInSubComment(recipientID string, sponsorID string, postID string, commentID string, subCommentID string, title string, commentContent string, subCommentContent string, TimeStamp int64) string {
	return s.r.InsertNotice(models.NoticeTypeCallInSubComment,recipientID,sponsorID,postID,commentID,subCommentID,title,commentContent,subCommentContent,TimeStamp)
}

func (s *noticeService) LikeNotice(recipientID, sponsorID, postID,commentID string, title string, commentContent string, TimeStamp int64) string {
	return s.r.InsertNotice(models.NoticeTypeLike,recipientID,sponsorID,postID,commentID,"",title,commentContent,"",TimeStamp)
}

func (s *noticeService) DislikeNotice(recipientID, sponsorID, postID,commentID string, title string, commentContent string, TimeStamp int64) string {
	return s.r.InsertNotice(models.NoticeTypeDislike,recipientID,sponsorID,postID,commentID,"",title,commentContent,"",TimeStamp)
}








