package main

import (
	"fmt"

	"github.com/go-pg/pg"
)

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
