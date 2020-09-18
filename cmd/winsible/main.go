package main

import (
	"flag"

	"github.com/abergmeier/winsible/internal/executor"
)

var (
	bucket = flag.String("bucket", "", "Bucket to download files from")
)

func main() {
	flag.Parse()
	config := mustReadConfig(*bucket)
	executor.MustRun(config)
}
