package mongo

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
)

var session *mgo.Session

const (
	url = "localhost"
)

// Data is result structure
type Data struct {
	_id   int
	Point int
}

func init() {
	f, err := os.OpenFile("/tmp/test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("error opening file :", err.Error())
	}

	log.SetOutput(f)

	session, err = mgo.Dial(url)

	if err != nil {
		log.Fatal("error connect mongo:", err.Error())
	}

	session.SetMode(mgo.Monotonic, true)
}

// Write function
func Write(id int, point int) {
	col := session.DB("test").C("test")

	_, err := col.Upsert(map[string]int{"_id": id}, bson.M{"_id": id, "point": point})

	if err != nil {
		log.Fatal(err.Error())
	}
}

// Read function
func Read(id int) {
	col := session.DB("test").C("test")

	result := Data{}
	err := col.Find(bson.M{"_id": id}).One(&result)

	if err != nil {
		log.Fatal("read error :", err.Error())
	}

	log.Printf("result : {_id : %d, point : %d}", id, result.Point)
}

// Remove all docs
func Remove() {
	col := session.DB("test").C("test")

	err := col.DropCollection()

	if err != nil {
		log.Fatal(err.Error())
	}
}

// IsEmpty function is check empty collection
func IsEmpty() bool {
	col := session.DB("test").C("test")

	n, err := col.Count()

	if err != nil {
		log.Fatal(err.Error())
	}

	return n == 0
}
