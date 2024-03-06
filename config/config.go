package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbName   = "challenge_goapi"
)

var DB *sql.DB

func Dbconnect() (db *sql.DB) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable dbname=%s", host, port, user, password, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	DB = db
	return
}
