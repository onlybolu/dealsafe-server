package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
	 "os"
	"dealsafe/database/sqlc"
)

var Queries *dealsafe.Queries

func ConnectDB() *sql.DB {

	DB_URL := os.Getenv("DB_STRING")

	db, err := sql.Open("postgres", DB_URL)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	fmt.Println("Connected to database")

	
	Queries = dealsafe.New(db)
	
	return db
}