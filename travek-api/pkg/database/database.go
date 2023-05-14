package database

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

const (
	db_host = "localhost:3306"
	db_name = "travek_db"
	db_user = "root"
	db_pass = "qwer"
)

func Get_db() *sql.DB {
	cnf := mysql.Config{
		User:   db_user,
		Passwd: db_pass,
		Net:    "tcp",
		Addr:   db_host,
		DBName: db_name,
	}
	db, err := sql.Open("mysql", cnf.FormatDSN())
	if err != nil {
		log.Fatal(err.Error())
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr.Error())
	}
	return db
}
