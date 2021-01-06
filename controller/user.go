package controller

import (
	"Geeklanet/models"
	"Geeklanet/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"time"
)


type UserController struct {
	S service.Service
}

func (c *UserController)PostSignin(context iris.Context) mvc.Result {

	receive := struct {
		UserName		string		`json:"userName"`
		UserPassword	string		`json:"userPassword"`
	}{}

	context.ReadJSON(&receive)


	if !c.S.User.CheckUserName(receive.UserName){
		return mvc.Response{
			Object: map[string]interface{}{
				"State":"no such user",
			},
		}
	}

	var state string
	session :=sessions.Get(context)
	if receive.UserPassword == c.S.User.GetUserPasswordByName(receive.UserName){
		userID := c.S.User.GetUserIDByName(receive.UserName)
		session.Set("userID",userID)
		session.Set("authenticated",true)
		state = "success"
	} else {
		state = "password wrong"
	}
	return mvc.Response{
		Object: map[string]interface{}{
			"State":state,
		},
	}
}

func (c *UserController)PostSignup(context iris.Context) mvc.Result {
	receive := struct {
		UserName		string		`json:"userName"`
		UserPassword	string	`json:"userPassword"`
		UserEmail		string		`json:"userEmail"`
		Gender			bool		`json:"email"`
	}{}

	context.ReadJSON(&receive)

	return mvc.Response{
		Object: map[string]interface{}{
			"userID":c.S.User.CreateNewAccount(receive.UserName,receive.UserPassword,receive.UserEmail,receive.Gender),
		},
	}

}

func (c *UserController)GetSignout(context iris.Context) mvc.Result {
	sessions.Get(context).Set("authenticated", false)

	return mvc.Response{
		Object: map[string]interface{}{
			"State":"success",
		},
	}
}

func (c *UserController)GetInfo(context iris.Context) mvc.Result {
	userID :=sessions.Get(context).GetString("userID")

	return mvc.Response{
		Object: c.S.User.GetUserBaseInfo(userID),
	}
}

func (c *UserController)GetInfoBy(userID string,context iris.Context) mvc.Result {

	return mvc.Response{
		Object: c.S.User.GetUserBaseInfo(userID),
	}
}

func (c *UserController) GetFollowsinfo(context iris.Context) mvc.Result {
	userID :=sessions.Get(context).GetString("userID")
	followsID := c.S.User.GetFollowsID(userID)
	usersInfo := make([]models.User,0)
	for _,followID := range followsID{
		userInfo := c.S.User.GetUserBaseInfo(followID)
		usersInfo = append(usersInfo,userInfo)
	}
	return mvc.Response{
		Object: map[string]interface{}{
			"Data":usersInfo,
		},
	}
}

func (c *UserController) GetFollowersinfo(context iris.Context) mvc.Result {
	userID :=sessions.Get(context).GetString("userID")
	followersID := c.S.User.GetFollowersID(userID)
	usersInfo := make([]models.User,0)
	for _,followerID := range followersID{
		userInfo := c.S.User.GetUserBaseInfo(followerID)
		usersInfo = append(usersInfo,userInfo)
	}
	return mvc.Response{
		Object: map[string]interface{}{
			"Data":usersInfo,
		},
	}
}

func (c *UserController) PostModifyName(context iris.Context) mvc.Result {
	receive := struct {
		UserName		string		`json:"userName"`
	}{}
	context.ReadJSON(&receive)
	userID := sessions.Get(context).GetString("userID")
	var state string
	if c.S.User.ModifyName(userID,receive.UserName){
		state = "success"
	} else {
		state = "failed"
	}
	return mvc.Response{
		Object: map[string]interface{}{
			"State":state,
		},
	}
}

func (c *UserController) PostModifyStatus(context iris.Context){
	receive := struct {
		Status		string		`json:"status"`
	}{}
	context.ReadJSON(&receive)
	userID := sessions.Get(context).GetString("userID")
	c.S.User.ModifyStatus(userID,receive.Status)
}

func (c *UserController) PostModifyAvatar(context iris.Context){
	receive := struct {
		Avatar		string		`json:"avatar"`
	}{}
	context.ReadJSON(&receive)
	userID := sessions.Get(context).GetString("userID")
	c.S.User.ModifyAvatar(userID,receive.Avatar)
}

func (c *UserController) PostFollowBy(followeeID string, context iris.Context) {
	userID := sessions.Get(context).GetString("userID")
	c.S.User.Follow(followeeID,userID)
	noticeID := c.S.Notice.Follow(userID,followeeID,time.Now().Unix())
	c.S.User.GetNotice(followeeID,noticeID)
}

func (c *UserController) DeleteFollowBy(followeeID string, context iris.Context) {
	userID := sessions.Get(context).GetString("userID")
	c.S.User.DeleteFollow(followeeID,userID)
	notice := c.S.Notice.DeleteFollow(userID,followeeID)
	c.S.User.DeleteNotice(followeeID,notice)
}