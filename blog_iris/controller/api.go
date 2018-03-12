package controller

import (
	"../models"

	"github.com/kataras/iris"
)
func Api(ctx iris.Context) {
	data := model.Getall(db)
	ctx.JSON(data)
}