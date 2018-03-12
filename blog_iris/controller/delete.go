package controller

import (
	"strconv"

	"../models"
	"github.com/kataras/iris"
)

func DeletePost(ctx iris.Context) {
	temp := ctx.Params().Get("id")
	i, _ := strconv.Atoi(temp)
	id := int16(i)
	model.Deletedata(db, id)
	ctx.HTML("<a href='/post' style='font-size: 18px; font-family: monospace'>/All Posts</a>")
	ctx.HTML("<h1>Deleted !!!</h1>")
}
