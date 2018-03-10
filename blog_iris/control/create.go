package main

import (
	"time" 
	"../Db"
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
func create(user,password,database string) {
	db := Db.ConnectToDb(user,password,database)

	app := iris.New()
	app.RegisterView(iris.HTML("./go_templates", ".html").Reload(true))

	app.Get("/create", func(ctx iris.Context) {
		if err := ctx.View("form.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
		}
	})
}