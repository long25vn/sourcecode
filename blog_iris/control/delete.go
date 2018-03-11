package control

import (
	"strconv"

	"../Db"
	"github.com/kataras/iris"
)

func DeletePost(ctx iris.Context) {
	temp := ctx.Params().Get("id")
	i, _ := strconv.Atoi(temp)
	i1 := int16(i)
	Db.Deletedata(db, i1)
	ctx.HTML("<a href='/post' style='font-size: 18px; font-family: monospace'>/All Posts</a>")
	ctx.HTML("<h1>Deleted !!!</h1>")
}
