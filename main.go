package main

import (
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
	var j []byte
	println("Hello, world from the RPI!")

	hue.GetBridgeInfo()

	a := raspi.NewAdaptor()
	s := gpio.NewPIRMotionDriver(a, "7")

	test := func() {
		s.On(gpio.MotionDetected, func(data interface{}) {
			println("Motion was detected!")
			j = []byte(`{"bri": 20}`)
			hue.SetLivingRoomBrightness(j)
		})
		s.On(gpio.MotionStopped, func(data interface{}) {
			println("Motion has stopped")
			j = []byte(`{"bri": 200}`)
			hue.SetLivingRoomBrightness(j)
		})
	}

	rpi := gobot.NewRobot("motionSensor", []gobot.Connection{a}, []gobot.Device{s}, test)

	if err := rpi.Start(); err != nil {
		log.Fatalf("Rpi could not start: " + err.Error())
	}
}
