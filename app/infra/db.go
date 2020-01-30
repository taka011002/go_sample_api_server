package infra

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"))

	var err error
	db, err = sql.Open("mysql", dbURI)

	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return  db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}