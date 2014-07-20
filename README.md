benchmark
=========

Benchmark database using go language

## Technology
* go v1.3
* redis server v2.8.12
* mysql server v5.6.19
* [zimbatm/direnv](https://github.com/zimbatm/direnv)
* [tools/godep](https://github.com/tools/godep)

## deps
* [garyburd/redigo](https://github.com/garyburd/redigo)
* [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

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
BenchmarkWrite     50000         73008 ns/op
BenchmarkRead      10000        124248 ns/op
ok      github.com/kosuda/benchmark/redis   5.731s
```

## mysql benchmark

```
> cd mysql
> go test -bench=".*"
BenchmarkWrite     10000        221457 ns/op
BenchmarkRead       1000       1911617 ns/op
ok      github.com/kosuda/benchmark/mysql   4.401s
```
