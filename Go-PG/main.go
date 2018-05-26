package main

func main() {
	db := Connect("postgres", "123456", "example", "localhost:5432")
	defer db.Close()
	initializeDatabase(db)

	selectinarray(db)
}
