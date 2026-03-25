package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"dealsafe/lib"

	_ "github.com/lib/pq"
)

func main() {
	lib.LoadEnv()
	dbURL := os.Getenv("DB_STRING")
	if dbURL == "" {
		log.Fatal("DB_STRING environment variable is not set")
	}

	fmt.Println("Attempting to connect to:", dbURL)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	fmt.Println("Successfully connected. Listing current columns of 'users' table...")
	
	rows, err := db.Query("SELECT column_name FROM information_schema.columns WHERE table_name = 'users'")
	if err != nil {
		log.Fatal("Error checking columns:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var colName string
		if err := rows.Scan(&colName); err != nil {
			log.Fatal(err)
		}
		fmt.Println("-", colName)
	}

	fmt.Println("Applying ALTER TABLE commands...")
	
	alterCommands := []string{
		"ALTER TABLE users ADD COLUMN IF NOT EXISTS first_name VARCHAR(255);",
		"ALTER TABLE users ADD COLUMN IF NOT EXISTS last_name VARCHAR(255);",
		"ALTER TABLE users ADD COLUMN IF NOT EXISTS phone_number VARCHAR(255);",
		"ALTER TABLE users ADD COLUMN IF NOT EXISTS address VARCHAR(255);",
		"ALTER TABLE users ADD COLUMN IF NOT EXISTS city VARCHAR(255);",
		"ALTER TABLE users ADD COLUMN IF NOT EXISTS company_name VARCHAR(255);",
		"ALTER TABLE users ADD COLUMN IF NOT EXISTS state VARCHAR(255);",
		"ALTER TABLE users ADD COLUMN IF NOT EXISTS zip_code VARCHAR(255);",
		"ALTER TABLE users ADD COLUMN IF NOT EXISTS terms_accepted BOOLEAN DEFAULT FALSE;",
		"ALTER TABLE users ADD COLUMN IF NOT EXISTS country VARCHAR(255);",
		"ALTER TABLE users ADD COLUMN IF NOT EXISTS test_pub_key VARCHAR(255);",
		"ALTER TABLE users ADD COLUMN IF NOT EXISTS test_priv_key VARCHAR(255);",
		"ALTER TABLE users ADD COLUMN IF NOT EXISTS live_pub_key VARCHAR(255);",
		"ALTER TABLE users ADD COLUMN IF NOT EXISTS live_priv_key VARCHAR(255);",
	}

	for _, cmd := range alterCommands {
		fmt.Printf("Executing: %s ", cmd)
		_, err = db.Exec(cmd)
		if err != nil {
			fmt.Printf(" -> ERROR: %v\n", err)
			log.Fatal(err)
		}
		fmt.Println(" -> Success")
	}

	fmt.Println("All ALTER commands completed. Re-checking columns...")
	rows2, _ := db.Query("SELECT column_name FROM information_schema.columns WHERE table_name = 'users'")
	for rows2.Next() {
		var colName string
		rows2.Scan(&colName)
		fmt.Println("- Ready:", colName)
	}
	rows2.Close()
}
