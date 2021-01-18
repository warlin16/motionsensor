package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config struct that will load external requirements for the app
type Config struct {
	Rpi struct {
		FirmataPort string `yaml:"firmataPort"`
	}
}

// FetchRpiDetails loads the configs for the rpi
func (c *Config) FetchRpiDetails() {
	yml, err := ioutil.ReadFile(os.Getenv("CONFIG_FILE"))
	if err != nil {
		log.Printf("An error ocurred reading the yml file: " + err.Error())
	}
	err = yaml.Unmarshal(yml, &c)
	if err != nil {
		log.Printf("Error unmarshalling yml: " + err.Error())
	}
}
