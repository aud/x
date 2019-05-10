package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
	"os"
)

type StorageClient struct {
	*storage.Client
}

var Storage StorageClient

const bucket = "x.dohm.dev"

func init() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		Log.Printf("%s", err)
		fmt.Fprintf(os.Stderr, "error: %v\n", err)

		os.Exit(1)
	}

	Storage = StorageClient{client}
}

func (client *StorageClient) Upload(filePath, object string, fs FileSystem) error {
	ctx := context.Background()

	file, err := fs.Open(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)

	wc.ObjectAttrs.ACL = append(
		wc.ObjectAttrs.ACL,
		storage.ACLRule{Role: storage.RoleReader, Entity: storage.AllUsers},
	)

	if _, err = io.Copy(wc, file); err != nil {
		return err
	}

	if err := wc.Close(); err != nil {
		return err
	}

	return nil
}
