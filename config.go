package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/QuestOJ/testdata-cache/typedef"
)

func readFile(filePath string) ([]byte, error) {
	Data, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	return Data, nil
}

func loadConfig(configFilePath string) error {
	config = typedef.Config{}

	configData, err := readFile(configFilePath)

	if err != nil {
		return err
	}

	err = json.Unmarshal(configData, &config)

	if err != nil {
		return err
	}

	return nil
}
