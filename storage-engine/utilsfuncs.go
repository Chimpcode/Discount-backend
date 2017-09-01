package storage

import (
	"os/exec"
	"log"
)

func GetUUID() string {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(out)
}
