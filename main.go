package main

import (
	"fmt"
	"log"
	"motionsensor/hue"
	"os"
	"path/filepath"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func init() {
	fileName, err := filepath.Abs("./config.yml")
	if err != nil {
		log.Printf("Failed to load config file: " + err.Error())
	}
	os.Setenv("CONFIG_FILE", fileName)
}

func main() {
	hue.GetBridgeInfo()

	a := raspi.NewAdaptor()
	s := gpio.NewPIRMotionDriver(a, "7")

	botAction := func() {
		s.On(gpio.MotionDetected, func(data interface{}) {
			toggleBrightness(true)
		})
		s.On(gpio.MotionStopped, func(data interface{}) {
			toggleBrightness(false)
		})
	}

	rpi := gobot.NewRobot("motionSensor", []gobot.Connection{a}, []gobot.Device{s}, botAction)

	if err := rpi.Start(); err != nil {
		log.Fatalf("Rpi could not start: " + err.Error())
	}
}

func toggleBrightness(t bool) {
	var brightness []byte
	var message string
	if t {
		brightness = []byte(`{"bri": 20}`)
		message = "Motion was detected!"
	} else {
		brightness = []byte(`{"bri": 200}`)
		message = "Motion has stopped"
	}
	hue.SetLivingRoomBrightness(brightness)
	fmt.Printf("%q\n", message)
}
