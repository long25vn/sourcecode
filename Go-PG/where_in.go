package main

import (
	"fmt"

	"github.com/go-pg/pg"
)

func wherein(db *pg.DB) {

	var title = []string{"comic", "science", "literature"}
	var book Book
	err := db.Model(&book).
		Where("category IN (?)", pg.Strings(title)).
		Select()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(book)
}
