package main

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Author struct
type Author struct {
	Id   int
	Name string
	Age  int
}

// Book struct
type Book struct {
	Id       int32
	Title    string
	Category string
	AuthorID int
	Author   Author
	Publish []int32 `pg:",array"`
}

func initializeDatabase(db *pg.DB) {
	var author Author
	var book Book

	for _, model := range []interface{}{&author, &book} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp:          false,
			FKConstraints: true,
			IfNotExists:   true,
		})
		if err != nil {
			panic(err)
		}
	}

}