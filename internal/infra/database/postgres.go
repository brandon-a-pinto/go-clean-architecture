package database

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func Start(driver, user, pass, host, port, name string) *sql.DB {
	db, err := sql.Open(driver, fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", driver, user, pass, host, port, name))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DB = db

	return db
}
