package main

import (
	"flag"
	"fmt"
	"os"
)

func version() {
	fmt.Println("Testdata Cache 1.0.0")
}

var port string
var dataDir string

func main() {
	version()

	flag.StringVar(&dataDir, "d", "./data", "data dir")
	flag.Parse()

	port = flag.Arg(0)

	if len(flag.Args()) == 0 || port == "" {
		fmt.Println("Usage: testdata-cache [port] [-d data dir] \n\nOptions:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	Init()
}
