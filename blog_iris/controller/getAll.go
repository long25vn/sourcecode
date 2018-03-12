package controller

import (
	"../models"
	"github.com/kataras/iris"
) 

func GetAllPost(ctx iris.Context) {
	data := model.Getall(db)
	ctx.View("select1.html", data)
}
