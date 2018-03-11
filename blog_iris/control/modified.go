package control

import (
	"strconv"

	"../Db"
	"github.com/kataras/iris"
)

func ModifiedPost(ctx iris.Context) {
	temp := ctx.Params().Get("id")
	i, _ := strconv.Atoi(temp)
	i1 := int16(i)
	data := Db.Modified(db, i1)
	ctx.View("modified.html", data)

}
