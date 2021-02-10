package main

import (
	"fmt"
	"log"
	"motionsensor/config"
	"motionsensor/hue"
	"os"
	"path/filepath"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

var c config.Config = config.Config{}

func init() {
	fileName, err := filepath.Abs("./config.yml")
	if err != nil {
		log.Printf("Failed to load config file: " + err.Error())
	}
	os.Setenv("CONFIG_FILE", fileName)
	os.Setenv("SPOTIFY_API_URL", "	https://api.spotify.com/v1/")
}

func main() {
	hue.GetBridgeInfo(c)

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
