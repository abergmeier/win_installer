package main

import (
	"flag"
	"os"

	"github.com/abergmeier/winsible/internal/config"
	"github.com/abergmeier/winsible/internal/executor"
)

var (
	file = flag.String("tasks", getCwd()+"/tasks.yaml", "YAML file to read tasks from")
)

func getCwd() string {
	d, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return d
}

func main() {
	flag.Parse()
	config := config.ReadYamlFile(*file)
	executor.MustRun(config)
}
