package main

import (
	"github.com/kosuda/benchmark/redis"
)

func main() {
	redis.Clear()
	redis.Write(10)
	redis.Read(10)
}