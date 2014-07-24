benchmark
=========
[![wercker status](https://app.wercker.com/status/4c82a8af8c2a95e21c589fdc380a7adb/m "wercker status")](https://app.wercker.com/project/bykey/4c82a8af8c2a95e21c589fdc380a7adb)

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

> godep restore

> ls
Godeps/    README.md  main.go    redis/     _vendor/

```
