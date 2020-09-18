package main

import (
	"log"
	"os"

	"github.com/abergmeier/winsible/internal/config"
)

func mustReadConfig() []interface{} {

	f, err := os.Open("tasks.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	return config.ReadYaml(f)
}
