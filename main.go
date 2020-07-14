package main

import (
	"flag"
	"fmt"
	"os"
)

const VERSION = "Testdata-Cache 1.2"

var port string
var dataDir string

func main() {
	fmt.Println(VERSION)

	flag.StringVar(&dataDir, "d", "./data", "data dir")
	flag.StringVar(&port, "p", "8081", "port")

	flag.Parse()

	if port == "" || dataDir == "" {
		fmt.Println("Usage: testdata-cache [-d data dir] [-p port] \n\nOptions:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	Init()
}
