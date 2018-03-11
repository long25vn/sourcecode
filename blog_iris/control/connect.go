package control

import (
	"encoding/json"
	"os"

	"../Db"
)

type Data struct {
	Database struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
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
	info := LoadfileJson("./Db/database.json")
	db = Db.ConnectToDb(info.Database.User, info.Database.Password, info.Database.Database)
}
