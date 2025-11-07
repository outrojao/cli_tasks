package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/dotenv-org/godotenvvault"
	_ "github.com/lib/pq"
)

func InitDatabase(status chan<- bool) {
	if err := godotenvvault.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	db := ConectDatabase()
	if db == nil {
		status <- false
		return
	}
	status <- true
	defer db.Close()
}

func ConectDatabase() (db *sql.DB) {
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	sslmode := os.Getenv("DB_SSLMODE")

	connection := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", user, dbname, password, host, sslmode)
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return
}
