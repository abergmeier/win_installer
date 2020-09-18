package gcstorage

import (
	"context"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

func Run(config map[string]interface{}) {
	bucketName := config["bucket"].(string)
	objectPath := config["object"].(string)
	dest := config["dest"].(string)
	mode := config["mode"].(string)

	if mode != "get" {
		log.Fatalf("Unknown Mode: %s", mode)
	}

	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)

	obj := bucket.Object(objectPath)
	rdr, err := obj.NewReader(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer rdr.Close()

	w, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	_, err = io.Copy(w, rdr)
	if err != nil {
		log.Fatal(err)
	}

	w.Sync()
}
