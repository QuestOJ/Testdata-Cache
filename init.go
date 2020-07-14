package main

import (
	"os"
	"path/filepath"

	"github.com/QuestOJ/testdata-cache/typedef"
)

var configFile string
var realPath string

var config typedef.Config

func Init() {
	realPath, _ = filepath.Abs(dataDir)

	configFile = dataDir + "/config.json"

	os.Mkdir(dataDir+"/testdata", 0770)

	log(2, "Service starting...")
	log(2, "Data path "+realPath)

	err := loadConfig(configFile)

	if err != nil {
		log(1, err.Error())
	}

	log(2, "Parse config.json success")

	startServer()
}
