package main

import (
	"flag"

	"github.com/abergmeier/winsible/internal/executor"
)

func main() {
	flag.Parse()
	config := mustReadConfig()
	executor.MustRun(config)
}
