package global


import (
	"io/ioutil"
	"encoding/json"
	"log"
	"github.com/alexflint/go-arg"
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

func ParseArguments() *Args{
	var args Args
	args.Conf = "server/cupointconfig.json"
	args.DP = 0
	arg.MustParse(&args)
	return &args
}
