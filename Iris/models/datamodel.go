package model

import (
)

type Post struct {
	Id             int16
	Title          string
	Alias          string
	Intro_text     string
	Full_text      string
	Image          string
	Published      string
	Published_at   string
	Categories     string
	Type           string
	Created_at     string
	Created_by     string
	Modified_at    string
	Modified_by    string
	Author_visible string
}

type Data struct {
	Database struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"database"`
}