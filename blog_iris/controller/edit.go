package controller

import (
	"strconv"
	"time"
	"../models"
	"github.com/kataras/iris"
)

func Edit(ctx iris.Context) {
	temp := ctx.Params().Get("id")
	i, _ := strconv.Atoi(temp)
	i1 := int16(i)
	data := model.Edit(db, i1)
	ctx.View("edit.html", data)

}
func Edited(ctx iris.Context) {
	post := model.Post{}
	err := ctx.ReadForm(&post)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
	ctx.ViewData("Title", post.Title)
	ctx.ViewData("Alias", post.Alias)
	ctx.ViewData("Image", post.Image)
	t1 := time.Now().Local().Format("2006-01-02 10:20:58")
	ctx.ViewData("Created_at", t1)
	ctx.ViewData("Published", post.Published)
	ctx.ViewData("Intro_text", post.Intro_text)
	ctx.ViewData("Full_text", post.Full_text)
	ctx.View("edited.html")
	temp := ctx.Params().Get("id")
	i, _ := strconv.Atoi(temp)
	i1 := int16(i)
	model.Update(db, post.Title, post.Alias, post.Intro_text, post.Full_text, post.Image, post.Published, t1, i1)

}