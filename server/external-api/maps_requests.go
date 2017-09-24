package external_api

import (
	"github.com/kataras/iris"
	"../../db-controller"
	"strconv"
)

func SetMapsAPI(api iris.Party)  {
	api.Get("/m/nearby", func(c iris.Context) {
		longitude, err := strconv.ParseFloat(c.URLParam("lon"), 64)
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{"error": "Invalid longitude value"})
			return
		}

		latitude, err := strconv.ParseFloat(c.URLParam("lat"), 64)
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{"error": "Invalid latitude value"})
			return
		}

		radio, err := strconv.ParseFloat(c.URLParam("radio"), 64)
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{"error": "Invalid radio value"})
			return
		}

		posts, err := db.GetAllPosts()
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{"error": err.Error()})
			return
		}

		nearbyPosts := make([]string, 0)
		for key, post := range posts {
			if latitude - radio <= post.Location.Latitude  &&  post.Location.Latitude <= latitude + radio {
				if longitude + radio <= post.Location.Longitude &&  post.Location.Longitude <= longitude + radio {
					nearbyPosts = append(nearbyPosts, key)
				}
			}
		}

		c.StatusCode(200)
		c.JSON(nearbyPosts)

	})
}

