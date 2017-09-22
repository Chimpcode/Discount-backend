package db

import "github.com/asdine/storm"

var Db *storm.DB


func checkError(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func Init() {
	var err error
	Db, err = storm.Open("my.db")
	checkError(err)
}

