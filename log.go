package main

import (
	"fmt"
	"os"
	"time"
)

var logDir string
var logFile string

func prepareLogDir() {
	logDir = dataDir + "/log/" + time.Now().Format("20060102")
	logFile = logDir + "/main.log"

	os.Mkdir(logDir, 0770)
}

func writeLogIntoFile(text string) {
	d1 := []byte(text)

	fl, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	defer fl.Close()
	_, err = fl.Write(d1)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

func log(level int, text string) {
	prepareLogDir()

	writeLogIntoFile(time.Now().Format("2006-01-02 15:04:05"))

	if level == 1 {
		writeLogIntoFile(" [Error] " + text + "\n")
		fmt.Println("[Error] " + text)
		os.Exit(2)
	} else if level == 2 {
		writeLogIntoFile(" [Warning] " + text + "\n")
		fmt.Println("[Warning] " + text)
	} else {
		writeLogIntoFile(" [Info] " + text + "\n")
		fmt.Println("[Info] " + text)
	}
}
