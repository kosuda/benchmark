package redis

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"math"
	"math/rand"
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
func Read(max int) {
	for i := 0; i < max; i++ {
		score, err := redis.Float64(conn.Do("ZSCORE", "rank", i))

		if err != nil {
			log.Fatal(err.Error())
		}

		count, err := redis.Int(conn.Do("ZCOUNT", "rank", score+1.0, math.Inf(0)))

		if err != nil {
			log.Fatal(err.Error())
		}

		log.Printf("(id, rank) = (%d, %d)", i, count+1)
	}
}

// Write function set records
func Write(max int) {
	for i := 0; i < max; i++ {
		value := rand.Intn(100000000000)
		_, err := conn.Do("ZADD", "rank", value, i)

		if err != nil {
			log.Fatal(err.Error())
		}

		log.Printf("key, value = (%d, %d)", i, value)
	}
}

// Clear function implements flushall command.
func Clear() {
	log.Println("fulushall")
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
