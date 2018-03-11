package main

import (
	_ "fmt"

	"./control"
	"github.com/kataras/iris"
)
const (
	user = "postgres"
	password = "123"
	database = "postgres"
)

func main() {
	control.ConnDb(user, password, database) 
 
	app := iris.New()
	app.RegisterView(iris.HTML("./templates", ".html").Reload(true))

	app.Get("/", func (ctx iris.Context) {
		ctx.View("menu.html")
	})
	app.Get("/post", control.GetAllPost)
	app.Get("/create", control.CreatePost)
	app.Post("/created", control.CreatedPost)
	app.Get("/preview", control.GetPublishPost)
	app.Get("/post/{id}",control.DetailsPost)
	app.Get("/post/modified/{id}", control.ModifiedPost)

	app.Run(iris.Addr(":8080"))
}
