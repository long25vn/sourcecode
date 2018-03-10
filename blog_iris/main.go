package main

import (
	_"fmt"
	"time"
	"./Db"
	"github.com/kataras/iris"
	"strconv"
) 

type Post struct {
	Id             int16
	Title          string
	Alias          string
	Intro_text     string
	Full_text      string
	Image          string
	Published      string
	Published_at   time.Time
	Categories     string
	Type           string
	Created_at     time.Time
	Created_by     string
	Modified_at    string
	Modified_by    string
	Author_visible string
}
func main() {
	user := "postgres"
	password := "123"
	database := "postgres"
	db := Db.ConnectToDb(user,password,database)

	app := iris.New()
	app.RegisterView(iris.HTML("./go_templates", ".html").Reload(true))

	app.Get("/create", func(ctx iris.Context) {
		if err := ctx.View("form.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
		}
	})


	app.Post("/created", func(ctx iris.Context) {
		post := Post{}
		err := ctx.ReadForm(&post)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
		}
		ctx.ViewData("Title", post.Title)
		ctx.ViewData("Alias", post.Alias)
		ctx.ViewData("Image", post.Image)
		t := time.Now()
		t1 := time.Now().Format(time.RFC850)
		ctx.ViewData("Created_at", t1)
		ctx.ViewData("Published", post.Published)
		ctx.ViewData("Full_text", post.Full_text)
		ctx.View("created.html")
		Db.Insertdata(db,post.Title,post.Alias,"null",post.Full_text,post.Image,post.Published,t)
		ctx.Writef("%v", post)

	})
	app.Get("/post/publish", func(ctx iris.Context) {
		data := Db.Getpublished(db)
		ctx.View("select1.html", data)

	})
	app.Get("/post/", func(ctx iris.Context) {
		data := Db.Getall(db)
		ctx.View("select1.html", data)

	})
	app.Get("/preview", func(ctx iris.Context) {
		data := Db.Getpublished(db)
		ctx.View("preview.html", data)

	})
	app.Get("/post/{id}", func(ctx iris.Context) {
		temp := ctx.Params().Get("id")
		i, _ := strconv.Atoi(temp)
		i1 := int16(i)
		data := Db.Modified(db, i1)
		ctx.View("details.html", data)

	})
	app.Get("/post/modified/{id}", func(ctx iris.Context) {
		temp := ctx.Params().Get("id")
		i, _ := strconv.Atoi(temp)
		i1 := int16(i)
		data := Db.Modified(db, i1)
		ctx.View("modified.html", data)

	})

	app.Run(iris.Addr(":8080"))
}