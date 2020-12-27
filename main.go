package main

import "github.com/kataras/iris/v12"

func main()  {
	app := iris.New()

	app.Get("/test", func(context iris.Context) {
		context.JSON(iris.Map{
			"test1":"Hello",
			"test2":"World",
		})
	})

	app.Run(iris.Addr(":8080"))
}
