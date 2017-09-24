package main

import (
	"github.com/kataras/iris"
	"../db-controller"
	"github.com/kataras/iris/middleware/recover"
	//"github.com/kataras/iris/middleware/logger"
	"../storage-engine"
	"../server/external-api"
	"log"
	"../global"
)

func feedDbWhitPosts() {
	log.Println("12 post creating")
	for i:=0;i<12;i++ {
		post := db.GetFakePost()
		uuid, err := db.SavePost(post)
		if err != nil {
			log.Println("error:", err.Error())
			return
		}
		log.Println(uuid)
	}

}




func main() {
	db.Init()

	conf := global.GetCupointConfigFromFile("server/cupointconfig.json")
	storage.InitStorage(conf)

	//feedDbWhitPosts()

	app := iris.New()
	app.Use(recover.New())

	tmpl := iris.HTML("./templates", ".html")
	tmpl.Reload(true) // reload templates on each request (development mode)
	app.RegisterView(tmpl)
	//app.RegisterView(iris.HTML("./templates", ".html").Layout("layout.html"))

	app.Get("/", func(c iris.Context) {
		c.Gzip(true)
		c.View("index.html")

	})

	api := app.Party("/api")

	external_api.SetPostsAPI(api)
	external_api.SetImagesAPI(api)

	app.Run(iris.Addr(":9300"), iris.WithoutServerError(iris.ErrServerClosed))
}


