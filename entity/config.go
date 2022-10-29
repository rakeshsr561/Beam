package entity

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Http struct {
		Retry             int    `json:"retry"`
		DestinationFolder string `json:"destination_folder"`
		SleepTIme         int    `json:"sleep_time"`
	} `json:"http"`
	Https struct {
		Retry             int    `json:"retry"`
		DestinationFolder string `json:"destination_folder"`
		SleepTIme         int    `json:"sleep_time"`
	} `json:"https"`
	Ftp struct {
		Retry             int    `json:"retry"`
		DestinationFolder string `json:"destination_folder"`
		SleepTIme         int    `json:"sleep_time"`
	} `json:"ftp"`
	Sftp struct {
		Retry             int    `json:"retry"`
		DestinationFolder string `json:"destination_folder"`
		SleepTIme         int    `json:"sleep_time"`
	} `json:"sftp"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
