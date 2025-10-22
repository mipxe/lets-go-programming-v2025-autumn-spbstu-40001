package main

import (
	"flag"

	"github.com/mipxe/task-3/internal/config"
)

func main() {
	configPath := flag.String("config", "", "path to yaml file")
	flag.Parse()

	if *configPath == "" {
		panic("config flag is required")
	}

	config, err := config.ReadConfig(*configPath)
	if err != nil {
		panic(err)
	}

}
