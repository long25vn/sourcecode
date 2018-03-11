package control

import (
	"strconv"

	"../Db"
	"github.com/kataras/iris"
)

func ApiDetails(ctx iris.Context) {
	temp := ctx.Params().Get("id")
	i, _ := strconv.Atoi(temp)
	id := int16(i)
	data := Db.Getbyid(db, id)
	ctx.JSON(data)
}
