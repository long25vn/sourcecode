package main

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	proto "./proto"
)

type Author struct {
	tableName struct{} `sql:"auth.author"`
	ID    int
	Name  string  `sql:",unique"`
}

var Db *pg.DB;

func InserData() {
	auth1 := Author {
		Name:"Cuong",
	}
	Db.Insert(&auth1)
}

func ConnectDb() error {
	Db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "123456",
		Database: "postgres",
		Addr:     "localhost:5432",
	})

	var n int
	_, err := Db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

func InitSchema() {
	var author Author
	var number proto.Shine

	// Tạo bảng
	for _, model := range []interface{}{&author, &number} {
		err := Db.CreateTable(model, &orm.CreateTableOptions{
			Temp:          false,
			FKConstraints: true,
			IfNotExists:   true,
		})
		if err != nil {
			panic(err)
		}
	}
}