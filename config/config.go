package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
)

const (
	filename = "test.json"
)

type logConf struct {
	Path string
}

type commonConf struct {
	Log logConf
}

type mysqlConf struct {
	User     string
	Password string
	Dbname   string
}

type psqlConf struct {
	User     string
	Password string
	Dbname   string
}

type sqlite3Conf struct {
	Dbpath string
}

type mongoConf struct {
	Host string
	Port int
}

type redisConf struct {
	Host string
	Port int
}

// Conf is configuration structure
type Conf struct {
	Common   commonConf
	Mysql    mysqlConf
	Postgres psqlConf
	Sqlite3  sqlite3Conf
	Mongo    mongoConf
	Redis    redisConf
}

// Configuration map
var Configuration Conf

func init() {
	_, selfname, _, _ := runtime.Caller(0)
	jsonStr, err := ioutil.ReadFile(filepath.Join(filepath.Dir(selfname), filename))

	if err != nil {
		log.Fatal(err.Error())
	}

	err = json.Unmarshal(jsonStr, &Configuration)

	if err != nil {
		log.Fatal(err.Error())
	}
}
