package repository

import (
	"Geeklanet/datasource"
	"Geeklanet/models"
	"go.mongodb.org/mongo-driver/bson"
)

type NoticeRepository struct {
	T datasource.Table
}

func convert2Notice(slice []interface{})[]models.Notice{
	var notices []models.Notice
	for _,e := range slice{
		var notice models.Notice
		bsonBytes, _ := bson.Marshal(e)
		bson.Unmarshal(bsonBytes, &notice)
		notices = append(notices,notice)
	}
	return notices
}

func (r NoticeRepository) QueryNoticeIDByRecipientIDandType(noticeType models.NoticeType, recipientID string) []models.Notice {
	return convert2Notice(r.T.Query(map[string]interface{}{
		"type":noticeType,
		"recipientID":recipientID,
	}))
}

func (r NoticeRepository) QueryNoticeIDBySponsorIDandCommentIDandType(noticeType models.NoticeType, sponsorID,commentID string) string {
	res := convert2Notice(r.T.Query(map[string]interface{}{
		"type":noticeType,
		"sponsorID":sponsorID,
		"commentID":commentID,
	}))

	if len(res)==0{
		return ""
	}else {
		return res[0].ID
	}
}

func (r NoticeRepository) QueryNoticeIDBySponsorIDandRecipientIDandType(noticeType models.NoticeType, sponsorID,recipientID string) string {
	res := convert2Notice(r.T.Query(map[string]interface{}{
		"type":noticeType,
		"sponsorID":sponsorID,
		"recipientID":recipientID,
	}))

	if len(res)==0{
		return ""
	}else {
		return res[0].ID
	}
}

func (r NoticeRepository) DeleteNoticeByID(noticeID string) {
	r.T.Delete(map[string]interface{}{
		"_id":noticeID,
	})
}

func (r NoticeRepository) DeleteNoticeByPostID(postID string) []models.Notice{
	return convert2Notice(r.T.Delete(map[string]interface{}{
		"postID":postID,
	}))
}

func (r NoticeRepository) DeleteNoticeByCommentID(commentID string) []models.Notice{
	return convert2Notice(r.T.Delete(map[string]interface{}{
		"commentID":commentID,
	}))
}

func (r NoticeRepository) DeleteNoticeBySubCommentID(subCommentID string) []models.Notice{
	return convert2Notice(r.T.Delete(map[string]interface{}{
		"subCommentID":subCommentID,
	}))
}

func (r NoticeRepository) QueryRecipientIDByNoticeID(noticeID string) string {
	res := convert2Notice(r.T.Query(map[string]interface{}{
		"_id":noticeID,
	}))

	if len(res)==0{
		return ""
	}else {
		return res[0].RecipientID
	}
}

func (r NoticeRepository) QueryNoticeByRecipientIDandUnread(recipientID string) []models.Notice {
	return convert2Notice(r.T.Query(map[string]interface{}{
		"recipientID":recipientID,
		"hasRead":false,
	}))
}

func (r NoticeRepository) InsertNotice(noticeType models.NoticeType, recipientID,sponsorID,postID,commentID,subCommentID,postTitle,commentContent,subCommentContent string,TimeStamp int64) string {
	return r.T.Insert(models.Notice{
		Type:              noticeType,
		RecipientID:       recipientID,
		SponsorID:         sponsorID,
		PostID:            postID,
		CommentID:         commentID,
		SubCommentID:      subCommentID,
		PostTitle:         postTitle,
		CommentContent:    commentContent,
		SubCommentContent: subCommentContent,
		HasRead:           false,
		TimeStampUnix:     TimeStamp,
	})
}