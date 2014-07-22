benchmark
=========

Benchmark database using go language

## Technology
* go v1.3
* redis server v2.8.12
* mysql server v5.6.19
* postgresql v9.3.4
* sqlite3 v3.7.13
* [zimbatm/direnv](https://github.com/zimbatm/direnv)
* [tools/godep](https://github.com/tools/godep)

## deps
* [garyburd/redigo](https://github.com/garyburd/redigo)
* [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
* [lib/pq](https://github.com/lib/pq)
* [mattn/go-sqlite3]( https://github.com/mattn/go-sqlite3)

## install
```
> go get github.com/kosuda/benchmark (OR ghq get kosuda/benchmark)

> cd $GOPATH/src/github.com/kosuda/benchmark (OR ghq look benchmark)

> godep get

> ls
Godeps/    README.md  main.go    redis/     vendor/

```

## redis benchmark

```
> cd redis
> go test -bench=".*"
PASS
BenchmarkWrite	   50000	     74234 ns/op
BenchmarkRead	     10000	    122150 ns/op
ok  	github.com/kosuda/benchmark/redis	5.761s
```

## mysql benchmark

```
> cd mysql
> go test -bench=".*"
PASS
BenchmarkWrite	   10000	    225284 ns/op
BenchmarkRead	      1000	   1950978 ns/op
ok  	github.com/kosuda/benchmark/mysql	4.480s
```

## postgres benchmark

```
> cd postgres
> go test -bench=".*"
PASS
BenchmarkWrite	    5000	    425098 ns/op
BenchmarkRead	     10000	    230049 ns/op
ok  	github.com/kosuda/benchmark/postgres	4.581s
```

## sqlite3 benchmark

```
> cd sqlite3
> go test -bench=".*"
PASS
BenchmarkWrite	     500	   5059931 ns/op
BenchmarkRead	     20000	     77313 ns/op
ok  	github.com/kosuda/benchmark/sqlite3	5.381s
```

## mongo benchmark

```
> cd mongo
> go test -bench=".*"
PASS
BenchmarkWrite	   10000	    136282 ns/op
BenchmarkRead	   10000	    130698 ns/op
ok  	github.com/kosuda/benchmark/mongo	2.734s
```
