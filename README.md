benchmark
=========

Benchmark database using go language

## Technology
* go v1.3
* redis server v2.8.12
* [zimbatm/direnv](https://github.com/zimbatm/direnv)
* [tools/godep](https://github.com/tools/godep)

## deps
* [garyburd/redigo](https://github.com/garyburd/redigo)

## install
```
> go get github.com/kosuda/benchmark OR ghq get kosuda/benchmark

> cd $GOPATH/src/github.com/kosuda/benchmark

> godep get

> ls
Godeps/    README.md  main.go    redis/     vendor/

```

## redis benchmark

```
> cd redis
> go test -bench=".*"
PASS
BenchmarkWrite	       1	5386683873 ns/op
BenchmarkRead	       1	9801904162 ns/op
ok  	github.com/kosuda/benchmark/redis	15.204s
```

## mysql benchmark
