package main

import (
	"os"
	"path/filepath"

	"github.com/QuestOJ/testdata-cache/typedef"
)

var logDir string
var logFile string
var configFile string
var realPath string

var config typedef.Config

func Init() {
	realPath, _ = filepath.Abs(dataDir)
	logDir = dataDir + "/log"
	logFile = logDir + "/main.log"
	configFile = dataDir + "/config.json"

	os.Mkdir(logDir, 0770)
	os.Mkdir(dataDir+"/testdata", 0770)

	log(3, "Service starting...")
	log(3, "Data path "+realPath)

	err := loadConfig(configFile)

	if err != nil {
		log(1, err.Error())
	}

	log(3, "Parse config.json success")

	startServer()
}
