package Db

import (
	"time"
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
