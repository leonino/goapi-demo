package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

func ConnectDb() (*sql.DB, error) {
	db, error := sql.Open("postgres", GetConnectionString())
	if error != nil {
		panic(error)
	}

	error = db.Ping()
	if error != nil {
		panic(error)
	}

	fmt.Println("Connected to database", GetDBName())
	return db, nil
}

func GetConnectionString() string {
	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", host, port, user, password, dbname)
}

func GetDBName() string {
	return dbname
}
