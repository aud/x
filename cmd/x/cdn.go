package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const basePath = "https://x.dohm.dev"

var (
	filePath, objectName string
	hidden               bool
)

func init() {
	flag.StringVar(&filePath, "path", "", "path to file")
	flag.StringVar(&objectName, "name", "", "object name")
	flag.BoolVar(&hidden, "hidden", false, "obfuscate url")

	flag.Parse()
}

func generateRandString(n int) (string, error) {
	rand.Seed(time.Now().UnixNano())

	bytes := make([]byte, n)
	_, err := rand.Read(bytes)

	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
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

	if objectName == "" && !hidden {
		var object = strings.Split(filePath, "/")
		objectName = object[len(object)-1]
	}

	if hidden {
		str, err := generateRandString(200)

		if err != nil {
			Log.Printf("error occurred during string generation %s", err)
			os.Exit(1)
		}

		objectName = str
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
