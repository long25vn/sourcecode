package main

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.com/kataras/iris"

	"github.com/kataras/iris/cache"
)


func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Get("/", cache.Handler(20*time.Second), routeDemo) // Cache in 20 minutes

	app.Run(iris.Addr(":8080"))
}

func routeDemo(ctx iris.Context) {
	fmt.Println("Handler executed. Content refreshed.")

	db := Connect("postgres", "123456", "postgres", "localhost:5432")
	defer db.Close()

	var book int32
	_, err := db.Query(&book, `
		SELECT count(id)
		FROM library.book WHERE category = 'literary'
	`)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(book)
	ctx.JSON(book)
}

func Connect(user, password, database, addr string) (db *pg.DB) {
	db = pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
		Addr:     addr,
	})
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		panic(err)
	}
	if n == 1 {
		fmt.Println("Connected")
	}
	return db
}
