package main

import (
	"fmt"

	"github.com/go-pg/pg"
)

func insertdb(db *pg.DB) {
	var author = Author{
		Name: "Nguyen Nhat Anh",
		Age:  63,
	}

	var book = Book{
		Title:    "Toi thay hoa vang tren co xanh",
		Category: "literature",
		AuthorID: 10,
		Publish:  []int32{2012, 2015},
	}
	err := db.Insert(&author)
	if err != nil {
		fmt.Print(err)
	}
	err = db.Insert(&book)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(book)
}
