package main

import (
	"log"
	"os"
	"path/filepath"
)

func init() {
	fileName, err := filepath.Abs("./config.yml")
	if err != nil {
		log.Printf("Failed to load config file: " + err.Error())
	}
	os.Setenv("CONFIG_FILE", fileName)
}

func main() {
	println("Hello, world from the RPI!")
}
