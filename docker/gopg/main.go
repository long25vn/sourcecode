package main

func main() {
	// time.Sleep(3 * time.Second)
	db := Connect("postgres", "123456", "postgres", "containerpostgres:5432")
	defer db.Close()
	initializeDatabase(db)

	selectinarray(db)
}
