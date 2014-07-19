package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestRemove(t *testing.T) {
	Remove()
}

func TestWrite(t *testing.T) {
	Write()
}

func TestRead(t *testing.T) {
	Read()
}
