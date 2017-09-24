package global


import (
	"io/ioutil"
	"encoding/json"
	"log"
)

func GetCupointConfigFromFile(path string) *CupointConf{
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var conf CupointConf
	err = json.Unmarshal(data, &conf)
	log.Println(conf)
	if err != nil {
		panic(err)
	}
	return &conf
}