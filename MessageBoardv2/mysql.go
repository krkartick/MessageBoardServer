package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql","root:kartick@tcp(192.168.56.5:3306)/MyDB")
	if err != nil {
		log.Fatal("Could not connect to MySql")
	}
	return db
}
