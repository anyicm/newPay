package main

import (
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"newPay/serve"
)

const defaultConfigPath = "./conf/conf.yaml"

func main() {
	var conf string
	pflag.StringVarP(&conf, "config", "c", defaultConfigPath, "config file path")
	pflag.Parse()
	f, err := ioutil.ReadFile(conf)
	if err != nil {
		log.Fatalf("read config failed, %s", err)
	}

	cfg := serve.NewConfiguration()
	err = yaml.Unmarshal(f, cfg)
	if err != nil {
		log.Fatalf("parse config failed, %s", err)
	}
	gate := serve.NewServer(cfg)
	gate.Run()
}
