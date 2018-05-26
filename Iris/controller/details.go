package controller

import (
	"strconv"

	"../models"
	"github.com/kataras/iris"
)

func DetailsPost(ctx iris.Context) {
	temp := ctx.Params().Get("id")
	i, _ := strconv.Atoi(temp)
	id := int16(i)
	data := model.Edit(db, id)
	ctx.View("details.html", data)

}
