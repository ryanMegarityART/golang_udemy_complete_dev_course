package main

import (
	"io"
	"os"
)

func main() {
	file := os.Args[1]
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, f)
}
