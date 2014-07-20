package redis

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"os"
)

var conn redis.Conn

const (
	address = "127.0.0.1:6379"
)

func init() {
	f, err := os.OpenFile("/tmp/test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("error opening file :", err.Error())
	}

	log.SetOutput(f)

	conn, err = redis.Dial("tcp", address)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Connected to redis store...")
}

// Read all records
func Read(id int) {
	point, err := conn.Do("GET", id)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("(id, point) = (%d, %d)", id, point)
}

// Write function set records
func Write(id int, point int) {
	_, err := conn.Do("SET", id, point)

	if err != nil {
		log.Fatal(err.Error())
	}
}

// Clear function implements flushall command.
func Clear() {
	_, err := conn.Do("FLUSHALL")

	if err != nil {
		log.Fatal(err.Error())
	}
}

// Empty check function
func Empty() bool {
	keys, err := redis.Strings(conn.Do("keys", "*"))

	if err != nil {
		log.Fatal(err.Error())
	}

	return len(keys) == 0
}
