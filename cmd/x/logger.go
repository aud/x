package main

import (
	"flag"
	"log"
	"os"
)

const (
	logPath = "cdn_debug.log"
)

var (
	Log *log.Logger
)

func init() {
	flag.Parse()

	file, err := os.Create(logPath)

	if err != nil {
		panic(err)
	}

	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}
