package main

import (
	"github.com/theoriz0/go-public/config"
	"github.com/theoriz0/go-public/flag"
	"github.com/theoriz0/go-public/log"
)

func main() {
	flag.HandleFlags()
	configFile := flag.GetConfigFile()
	config := config.ReadConfigFile(configFile)
	log.Init(log.LogOptions(config))
	log.Infow("Hello world")
}
