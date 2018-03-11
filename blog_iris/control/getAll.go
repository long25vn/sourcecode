package control

import (
	"../Db"
	"github.com/kataras/iris"
)

func GetAllPost(ctx iris.Context) {
	data := Db.Getall(db)
	ctx.View("select1.html", data)
}
