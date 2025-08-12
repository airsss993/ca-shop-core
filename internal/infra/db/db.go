package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln("failed to load .env file")
	}

	dsn := os.Getenv("DB_CONN")
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal("failed to connect to db", err)
	}

	err = db.Ping()

	return db
}
