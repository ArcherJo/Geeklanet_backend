package controller

import (
	"Geeklanet/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type NoticeController struct {
	S service.Service
}

func (c NoticeController) GetUnread(context iris.Context) mvc.Result {
	userID := sessions.Get(context).GetString("userID")
	return mvc.Response{
		Object: map[string]interface{}{
			"Data":c.S.Notice.GetUnreadNoticeByRecipientID(userID),
		},
	}
}



func (c NoticeController) GetFollow(context iris.Context) mvc.Result {
	userID := sessions.Get(context).GetString("userID")
	return mvc.Response{
		Object: map[string]interface{}{
			"Data":c.S.Notice.GetFollowNoticeByRecipientID(userID),
		},
	}
}



func (c NoticeController) GetCall(context iris.Context) mvc.Result {
	userID := sessions.Get(context).GetString("userID")
	return mvc.Response{
		Object: map[string]interface{}{
			"Data":c.S.Notice.GetCallNoticeByRecipientID(userID),
		},
	}
}



func (c NoticeController) GetRecommend(context iris.Context) mvc.Result {
	userID := sessions.Get(context).GetString("userID")
	return mvc.Response{
		Object: map[string]interface{}{
			"Data":c.S.Notice.GetUnreadNoticeByRecipientID(userID),
		},
	}
}


func (c NoticeController) GetLike(context iris.Context) mvc.Result {
	userID := sessions.Get(context).GetString("userID")
	return mvc.Response{
		Object: map[string]interface{}{
			"Data":c.S.Notice.GetLikeNoticeByRecipientID(userID),
		},
	}
}

func (c NoticeController) GetReply(context iris.Context) mvc.Result {
	userID := sessions.Get(context).GetString("userID")
	return mvc.Response{
		Object: map[string]interface{}{
			"Data":c.S.Notice.GetReplyNoticeByRecipientID(userID),
		},
	}
}

