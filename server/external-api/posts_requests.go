package external_api

import (
	"github.com/kataras/iris"
	"../../db-controller"
	"log"
	"time"
)
func SetPostsAPI (api iris.Party) {

	api.Get("/p/{id}", func(c iris.Context) {
		id := c.Params().Get("id")
		log.Println(id)
		post, err := db.GetPostById(id)
		if err!=nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{"error": err.Error()})
			return
		}
		c.StatusCode(iris.StatusOK)
		c.JSON(post)
	})

	api.Get("/p", func(c iris.Context) {
		posts, err := db.GetAllPosts()
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{"error": err.Error()})
			return
		}
		c.StatusCode(iris.StatusOK)
		c.JSON(posts)
	})

	api.Post("/p/new", func(c iris.Context) {
		var post db.Post
		err := c.ReadForm(&post)
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{"error": err.Error()})
			return
		}
		post.Id = db.GetUUID()
		post.CreatedAt = time.Now()
	})
}
