package main

import (
	"fmt"

	"github.com/go-pg/pg"
)

func selectinarray(db *pg.DB) {
	var book []Book
	err := db.Model(&book).
		Where("publish @> '{?}'", 2015).
		Select()

	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(book)
}
