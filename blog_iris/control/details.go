package control

import (
	"strconv"

	"../Db"
	"github.com/kataras/iris"
)

func DetailsPost(ctx iris.Context) {
	temp := ctx.Params().Get("id")
	i, _ := strconv.Atoi(temp)
	i1 := int16(i)
	data := Db.Edit(db, i1)
	ctx.View("details.html", data)

}
