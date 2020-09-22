package gcstorage

import (
	"context"
	"crypto/md5"
	"errors"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"github.com/abergmeier/winsible/internal/filehash"
)

func Run(config map[string]interface{}) error {
	bucketConfig, ok := config["bucket"]
	if !ok {
		return errors.New("Missing bucket config")
	}
	if bucketConfig == nil {
		return errors.New("Missing value in bucket config")
	}
	bucketName := bucketConfig.(string)
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
	attrs, err := obj.Attrs(ctx)
	if err == nil {
		downloaded, err := isDownloaded(dest, attrs.MD5)
		if err == nil {
			if downloaded {
				return nil // Already done
			}
		} else {
			log.Print("Could not check local MD5 - falling back to downloading")
		}
	} else {
		log.Print("Could not check remote MD5 - falling back to downloading")
	}

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

	return w.Sync()
}

func isDownloaded(dest string, md5Hash []byte) (bool, error) {
	r, err := os.Open(dest)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	defer r.Close()
	expectedHash := [md5.Size]byte{}
	copy(expectedHash[:], md5Hash)
	downloaded, err := filehash.ReaderHasMD5(r, expectedHash)
	return downloaded, err
}
