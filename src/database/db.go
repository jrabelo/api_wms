package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/godror/godror"
	"github.com/joho/godotenv"
)

func Connect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dns := os.Getenv("DB_USER") + "/" + os.Getenv("DB_PASSWD") + "@" + os.Getenv("DB_SERVER") + "/" + os.Getenv("DB_SID")
	connection, err := sql.Open("godror", dns)
	if err != nil {
		log.Fatal(err)
	}

	return connection
}
