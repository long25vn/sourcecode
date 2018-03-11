package control

import (
	"time"

	"../Db"
	"github.com/go-pg/pg"
	"github.com/kataras/iris"
)

var db *pg.DB

func CreatePost(ctx iris.Context) {
	if err := ctx.View("form.html"); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
}
func CreatedPost(ctx iris.Context) {
	post := Db.Post{}
	err := ctx.ReadForm(&post)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
	ctx.ViewData("Title", post.Title)
	ctx.ViewData("Alias", post.Alias)
	ctx.ViewData("Image", post.Image)
	t := time.Now()
	t1 := time.Now().Format(time.RFC850)
	ctx.ViewData("Created_at", t1)
	ctx.ViewData("Published", post.Published)
	ctx.ViewData("Full_text", post.Full_text)
	ctx.View("created.html")
	Db.Insertdata(db, post.Title, post.Alias, "null", post.Full_text, post.Image, post.Published, t)
	ctx.Writef("%v", post)

}
