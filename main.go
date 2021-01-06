package main

import (
	"Geeklanet/controller"
	"Geeklanet/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)


func main()  {
	app := iris.New()

	sess := sessions.New(sessions.Config{Cookie: "userID", AllowReclaim: true})
	app.Use(sess.Handler())


	app.Use(func(context *context.Context) {
		iris.New().Logger().Info("To:",context.Path())
		if context.Path()=="/user/signin" || context.Path()=="/user/signup"{
			context.Next()
		}
		//if auth, _ := sessions.Get(context).GetBoolean("authenticated"); auth{
		//	context.StatusCode(iris.StatusForbidden)
		//	return
		//}
		context.Next()
	})
	Service := service.NewService()
	mvc.New(app.Party("/user")).Register(*Service).Handle(new(controller.UserController))
	mvc.New(app.Party("/post")).Register(*Service).Handle(new(controller.PostController))
	mvc.New(app.Party("/notice")).Register(*Service).Handle(new(controller.NoticeController))
	mvc.New(app.Party("/tag")).Register(*Service).Handle(new(controller.TagController))
	mvc.New(app.Party("/achievement")).Register(*Service).Handle(new(controller.AchievementController))


	app.Run(iris.Addr(":8080"))
}
