package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		HomeController()
	})

	app.Run(iris.Addr("127.0.0.1:1234"))
}
