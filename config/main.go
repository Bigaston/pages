package config

import (
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type ConfigType struct {
	Global GlobalConfigType `toml:"global"`
}

type GlobalConfigType struct {
	BaseHostname string `toml:"base_hostname"`
}

var Config ConfigType

func ReadConfig(file string) {
	log.Println("Reading config from", file)

	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	err = toml.Unmarshal(content, &Config)

	if err != nil {
		panic(err)
	}
}
