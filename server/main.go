package main

import (
	"github.com/kataras/iris"
	"../db-controller"
	"github.com/kataras/iris/middleware/recover"
	//"github.com/kataras/iris/middleware/logger"
)


func main() {
	db.Init()
	app := iris.New()
	app.Use(recover.New())

	tmpl := iris.HTML("./templates", ".html")
	tmpl.Reload(true) // reload templates on each request (development mode)
	app.RegisterView(tmpl)
	//app.RegisterView(iris.HTML("./templates", ".html").Layout("layout.html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.Gzip(true)
		ctx.View("index.html")

	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}


