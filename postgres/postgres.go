package mysql

import (
	"database/sql"
	"log"
	"os"
	"strconv"
)

const (
	user     = "kosuda"
	password = ""
	database = "postgres"
)

var con *sql.DB

func init() {
	f, err := os.OpenFile("/tmp/test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		log.Fatal("error opening file :", err.Error())
	}
	log.SetOutput(f)

	//con, err = sql.Open("postgres", "user="+user+" password="+password+" dbname="+database)
	con, err = sql.Open("postgres", "postgres://"+user+":"+password+"@/"+database+"?sslmode=disable")

	if err != nil {
		log.Fatal(err.Error())
	}
}

// Write function insert record
func Write(id int, point int) {
	_, err := con.Exec("SELECT replace(" + strconv.Itoa(id) + "," + strconv.Itoa(point) + ")")

	if err != nil {
		log.Fatal(err.Error())
	}
}

// Remove function
func Remove() {
	_, err := con.Exec("DELETE FROM test")

	if err != nil {
		log.Fatal(err.Error())
	}
}

// Read function
func Read(id int) {
	rows, err := con.Query("SELECT point FROM test WHERE id=" + strconv.Itoa(id))

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		var point int

		if err := rows.Scan(&point); err != nil {
			log.Fatal(err.Error())
		}

		log.Printf("point = %d", point)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err.Error())
	}
}
