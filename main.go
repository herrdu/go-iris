package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("write text")
	})

	app.Run(iris.Addr(":8080"))
}
