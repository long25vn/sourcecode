package Db

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
)

type Post struct {
	Id             int16
	Title          string
	Alias          string
	Intro_text     string
	Full_text      string
	Image          string
	Published      string
	Published_at   time.Time
	Categories     string
	Type           string
	Created_at     time.Time
	Created_by     string
	Modified_at    string
	Modified_by    string
	Author_visible string
}
func ConnectToDb(user, password, database string) (db *pg.DB) {
	db = pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
	})
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		panic(err)
	}
	if n==1 {
		fmt.Println("Connected")
	}
	return db
}

func Getdata(db *pg.DB, id int16) (data Post) {
	data = Post{Id: id}
	err := db.Select(&data)
	if err != nil {
		panic(err)
	}
	return data
}
func Getall(db *pg.DB) (data []Post) {
	_ = db.Model(&data).Select()
	return data
}
func Getpublished(db *pg.DB) (data []Post) {
	_ = db.Model(&data).Where("published LIKE ?", "publish%").Select()
	return data
}
func Modified(db *pg.DB, idpost int16) (data Post) {
	data = Post{Id: idpost}
	_ = db.Select(&data)
	fmt.Println(data)
	return data
}
func Insertdata(db *pg.DB, title,alias,intro_text,full_text,image string, published string, created_at time.Time) (data Post) {
	data = Post{Title: title, Alias: alias, Intro_text: intro_text, 
		Full_text: full_text, Image: image, Published: published, Created_at: created_at}
	err := db.Insert(&data)
	if err != nil {
		panic(err)
	}
	return data
}
