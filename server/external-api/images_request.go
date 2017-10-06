package external_api

import (

	"io/ioutil"
	"fmt"
	"log"
	"github.com/kataras/iris"
	"../../storage-engine"
)

func SetImagesAPI(api iris.Party)  {
	api.Get("/i/{id}", func(c iris.Context) {
		id := c.Params().Get("id")
		log.Println("Getting image ", id)
		data, err := storage.GetImage(id)
		if err != nil {
			log.Println(err)
			c.JSON(iris.Map{"error": err.Error()})
		}
		err = ioutil.WriteFile("./tmp/img_tmp", data, 0666)//os.ModeAppend|os.ModeDir)
		if err != nil {
			log.Println(err)
			c.JSON(iris.Map{"error": err.Error()})
		}
		err = c.SendFile("./tmp/img_tmp", "image.jpg")
		if err != nil {
			log.Println(err)
			c.JSON(iris.Map{"error": err.Error()})
		}
	})

	api.Post("/i/{id}", iris.LimitRequestBodySize(10<<20), func(c iris.Context) {
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
			c.JSON(iris.Map{"error": err.Error()})
		}
	})
}