package control

import (
	"../Db"
)

func ConnDb(user, password, database string) {
	db = Db.ConnectToDb(user, password, database)
}
