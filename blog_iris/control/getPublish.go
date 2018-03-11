package control

import (
	"../Db"
	"github.com/kataras/iris"
) 

func GetPublishPost(ctx iris.Context) {
	data := Db.Getpublished(db)
	ctx.View("preview.html", data)
} 