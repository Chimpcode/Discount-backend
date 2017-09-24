package external_api

import (
	"github.com/kataras/iris"
)

func SetMapsAPI(api iris.Party)  {
	api.Get("/m/nearby", func(c iris.Context) {


	})
}

