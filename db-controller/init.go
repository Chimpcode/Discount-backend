package db

import "github.com/asdine/storm"

func checkError(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func Init() {
	db, err := storm.Open("my.db")
	checkError(err)
	defer db.Close()
}

