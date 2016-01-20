package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/naoina/toml"
)

const configFilename = "config"

var config Config

type Config struct {
	inputFile string
	Settings  struct {
		TranscactionSpacing int
	}
}

func LoadConfig() {
	f, err := os.Open(configFilename)
	defer f.Close()
	if os.IsNotExist(err) {
		setDefaultConfig(&config)
		return
	} else if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(f)

	if err = toml.Unmarshal(data, &config); err != nil {
		panic(err)
	}
}

func SaveConfig() {
	data, err := toml.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(configFilename)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	f.Write(data)
}

func setDefaultConfig(cfg *Config) {
	cfg.Settings.TranscactionSpacing = 1
}
