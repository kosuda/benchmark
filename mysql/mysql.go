package mysql

import (
	"database/sql"
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

// Remove function
func Remove() {
	_, err := con.Exec("DELETE FROM ranking WHERE id=?", 1)

	if err != nil {
		log.Fatal(err.Error())
	}
}

// Read function
func Read() {
	rows, err := con.Query("SELECT * FROM ranking")

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		var id int
		var point int

		if err := rows.Scan(&id, &point); err != nil {
			log.Fatal(err.Error())
		}

		log.Printf("%d is %d", id, point)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err.Error())
	}
}
