package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	user     = "kosuda"
	password = ""
	database = "test"
)

var con *sql.DB

func init() {
	var err error
	con, err = sql.Open("mysql", user+":"+password+"@/"+database)

	if err != nil {
		log.Fatal(err.Error())
	}
}

// Write function insert record
func Write() {
	_, err := con.Exec("INSERT INTO ranking (id,point) VALUES(?,?)", 1, 10)

	if err != nil {
		log.Fatal(err.Error())
	}
}
