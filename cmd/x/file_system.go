package main

import (
	"io"
	"os"
)

type FileSystem interface {
	Open(name string) (file, error)
	Stat(name string) (os.FileInfo, error)
}

type file interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	Stat() (os.FileInfo, error)
}

func (osFS) Open(name string) (file, error) {
	return os.Open(name)
}

func (osFS) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

type osFS struct{}

var fs FileSystem = osFS{}
