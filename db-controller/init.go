package db

import (
	"github.com/asdine/storm"
	"time"
	"github.com/boltdb/bolt"
)

var Db *storm.DB


func checkError(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func Init() {
	var err error
	Db, err = storm.Open("my.db", storm.BoltOptions(0644, &bolt.Options{Timeout: 1 * time.Second}))
	checkError(err)
}

