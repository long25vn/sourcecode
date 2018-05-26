package controller

import (
	"encoding/json"
	"os"

	"../models"
)

type Data struct {
	Database struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
		Addr string `json:"addr"`
	} `json:"database"`
}

func LoadfileJson(file string) Data {
	var data Data
	configFile, _ := os.Open(file)
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&data)
	return data
}

func ConnDb() {
	info := LoadfileJson("./models/database.json")
	db = model.ConnectToDb(info.Database.User, info.Database.Password, info.Database.Database, info.Database.Addr)
}
