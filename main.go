package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	remote := flag.Bool("remote", false, "use remote config")
	flag.Parse()

	if *remote {
		run(NewRemoteConfigReader("consul", "localhost:8500", "consul_config_demo/config.json"))
	} else {
		run(NewLocalConfigReader())
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}

func run(configReader ConfigReader) {
	err := configReader.ReadConfig()
	if err != nil {
		panic(err)
	}
}
