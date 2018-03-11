package control

import (
	"../Db"

	"github.com/kataras/iris"
)
func Api(ctx iris.Context) {
	data := Db.Getall(db)
	ctx.JSON(data)	
}