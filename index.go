package main

import (
	"goggles/controllers"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	//api endpoints
	app.Get("/movies", func(ctx iris.Context) {
		mv := controllers.MoviesController{Cntx: ctx}
		mv.Get()
	})

	app.Get("/movies/{id}", func(ctx iris.Context) {
		mv := controllers.MoviesController{Cntx: ctx}
		mv.GetByID(ctx.Params().Get("id"))
	})

	app.Get("/dashboard", func(ctx iris.Context) {
		(new(controllers.HomeController)).ShowDashboard()
	})

	app.Get("/", func(ctx iris.Context) {
		(new(controllers.HomeController)).Show()
	})

	app.Run(iris.Addr("127.0.0.1:1234"))
}
