package controller

import (
	"../models"
	"github.com/kataras/iris"
) 

func GetPublishPost(ctx iris.Context) {
	data := model.Getpublished(db)
	ctx.View("preview.html", data)
} 