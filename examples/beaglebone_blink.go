// +build example
//
// Do not build by default.

package main

import (
	"time"

	"github.com/matipan/gobot"
	"github.com/matipan/gobot/drivers/gpio"
	"github.com/matipan/gobot/platforms/beaglebone"
)

func main() {
	beagleboneAdaptor := beaglebone.NewAdaptor()
	led := gpio.NewLedDriver(beagleboneAdaptor, "P9_12")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{beagleboneAdaptor},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}
