package main

import (
	"github.com/kataras/iris"
	"../db-controller"
	"github.com/kataras/iris/middleware/recover"
	//"github.com/kataras/iris/middleware/logger"
	"log"
	"../storage-engine"
	"io/ioutil"
	"os"
	"fmt"
)


func main() {
	db.Init()
	storage.InitStorage()
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
	apiRoutes := app.Party("/api")

	apiRoutes.Get("/images/{id:string}", func(c iris.Context) {
		id := c.Params().Get("id")
		log.Println("Getting image from ", id)
		data, err := storage.GetImage(id)
		if err != nil {
			log.Println(err)
			c.Err()
		}
		err = ioutil.WriteFile("./tmp/imageTemp.jpg", data, os.ModeAppend|os.ModeDir)
		if err != nil {
			log.Println(err)
			c.Err()
		}
		err =  c.SendFile("./tmp/imageTemp.jpg", "image.jpg")
		if err != nil {
			log.Println(err)
			c.Err()
		}
	})

	apiRoutes.Post("/images/{id:string}", iris.LimitRequestBodySize(10<<20), func(c iris.Context) {
		id := c.Params().Get("id")
		log.Println("Uploading image from ", id)
		fmt.Println("-----0>", c.GetContentType())
		file, _, err := c.FormFile("file")
		fmt.Println("-----0", err)

		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		}
		defer file.Close()
		fmt.Println("-----1")

		err = storage.UploadImage(file, id)
		fmt.Println("-----2", err)

		if err != nil {
			log.Println(err.Error())
			fmt.Println("-----", err)
			c.Err()
		}
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}


