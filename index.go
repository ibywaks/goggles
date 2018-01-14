package main

import (
	"strconv"
	"goggles/controllers"
	"goggles/models"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	db, _ := gorm.Open("sqlite3", "./db/gorm.db")

	db.AutoMigrate(&models.Movies{}, &models.EndPoints{}, &models.EndPointsCall{})

	// defer db.Close()

	//api endpoints
	app.Get("/api/movies", func(ctx iris.Context) {
		mv := controllers.MoviesController{Cntx: ctx}
		mv.Get()
	})

	app.Post("/api/movies", func(ctx iris.Context) {
		mv := controllers.MoviesController{Cntx: ctx}
		mv.Add()
	})

	app.Get("/api/movies/{id}", func(ctx iris.Context) {
		ID,_ := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)
		mv := controllers.MoviesController{Cntx: ctx}
		mv.GetByID(ID)
	})

	app.Post("/api/movies/{id}", func(ctx iris.Context) {
		ID,_ := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)
		
		mv := controllers.MoviesController{Cntx: ctx}
		mv.Edit(ID)
	})

	app.Delete("/api/movies/{id}", func(ctx iris.Context) {
		ID,_ := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)
		
		mv := controllers.MoviesController{Cntx: ctx}
		mv.Delete(ID)
	})

	app.Get("/admin/dashboard", func(ctx iris.Context) {
		(new(controllers.HomeController)).ShowDashboard()
	})

	app.Get("/", func(ctx iris.Context) {
		(new(controllers.HomeController)).Show()
	})

	app.Run(iris.Addr("127.0.0.1:1234"))
}
