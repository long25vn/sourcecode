package controller

import (
	"strconv"

	"../models"
	"github.com/kataras/iris"
)

func ApiDetails(ctx iris.Context) {
	temp := ctx.Params().Get("id")
	i, _ := strconv.Atoi(temp)
	id := int16(i)
	data := model.Getbyid(db, id)
	ctx.JSON(data)
}
