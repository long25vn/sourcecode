package main

import (
	"fmt"
	"time"

	"./controller"
	"github.com/kataras/iris"
)

func main() {
	fmt.Println(time.Now().Local().Format("2006-01-02"))
	controller.ConnDb() 

	app := iris.New()
	app.RegisterView(iris.HTML("./templates", ".html").Reload(true))

	app.Get("/", func(ctx iris.Context) {
		ctx.View("menu.html")
	})
	app.Get("/post", controller.GetAllPost)
	app.Get("/create", controller.CreatePost)
	app.Post("/created", controller.CreatedPost)
	app.Get("/preview", controller.GetPublishPost)
	app.Get("/post/{id}", controller.DetailsPost)
	app.Get("/post/edit/{id}", controller.Edit)
	app.Post("/post/edited/{id}", controller.Edited)
	app.Get("/post/delete/{id}", controller.DeletePost)

	app.Get("/api", controller.Api)
	app.Get("/api/{id}", controller.ApiDetails)

	app.Run(iris.Addr(":8080"))
}
