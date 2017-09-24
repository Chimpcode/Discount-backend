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

	"time"
	"math/rand"
)

func feedDbWhitPosts(howMuch int) {

	err := db.DeleteAllPosts()
	if err != nil {
		log.Println("error:", err.Error())
		return
	}

	log.Println(howMuch, "post will created")

	boundA := db.Location{Longitude: -12.007243, Latitude: -77.109899}
	boundB := db.Location{Longitude: -12.087220, Latitude: -76.959972}

	for i:=0;i<howMuch;i++ {
		post := db.GetFakePost(boundA,boundB)
		uuid, err := db.SavePost(post)
		if err != nil {
			log.Println("error:", err.Error())
			return
		}
		log.Println(uuid)
	}

}




func main() {
	rand.Seed(time.Now().Unix())

	db.Init()

	args := global.ParseArguments()

	configPath := args.Conf
	conf := global.GetCupointConfigFromFile(configPath)
	storage.InitStorage(conf)

	feedDbWhitPosts(args.DP)

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


