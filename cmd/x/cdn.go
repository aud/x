package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const basePath = "https://x.dohm.dev"

var (
	filePath, objectName string
)

func init() {
	flag.StringVar(&filePath, "path", "", "path to file")
	flag.StringVar(&objectName, "name", "", "object name")

	flag.Parse()
}

func main() {
	if filePath == "" {
		Log.Printf("%s: %s", "invalid file path", filePath)
		fmt.Fprintf(os.Stderr, "invalid file path. missing valid `-path` argument\n")

		os.Exit(1)
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		Log.Printf("%s", err)
		fmt.Fprintf(os.Stderr, "error: %v\n", err)

		os.Exit(1)
	}

	if objectName == "" {
		var object = strings.Split(filePath, "/")
		objectName = object[len(object)-1]
	}

	Log.Printf("%s: %s, %s: %s", "uploading file", filePath, "object", objectName)

	if err := Storage.Upload(filePath, objectName, fs); err != nil {
		Log.Printf("%s", err)
		fmt.Fprintf(os.Stderr, "error occurred during upload: %v\n", err)

		os.Exit(1)
	}

	Log.Printf("%s: %s", "successfully uploaded file", filePath)

	fmt.Printf("%s/%s\n", basePath, objectName)
}
