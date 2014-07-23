package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/kosuda/benchmark/mongo"
	_ "github.com/kosuda/benchmark/redis"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {}
