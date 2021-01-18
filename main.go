package main

import (
	"log"
	"motionsensor/config"
	"os"
	"path/filepath"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
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
	c := config.Config{}
	c.FetchRpiDetails()
	a := firmata.NewAdaptor(c.Rpi.FirmataPort)
	s := gpio.NewPIRMotionDriver(a, "5")

	test := func() {
		s.On(gpio.MotionDetected, func(data interface{}) {
			println("Motion was detected!")
		})
		s.On(gpio.MotionStopped, func(data interface{}) {
			println("Motion has stopped")
		})
	}

	rpi := gobot.NewRobot("motionSensor", []gobot.Connection{a}, []gobot.Device{s}, test)

	if err := rpi.Start(); err != nil {
		log.Fatalf("Rpi could not start: " + err.Error())
	}
}
