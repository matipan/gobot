/*
Package particle provides the Gobot adaptor for the Particle Photon and Electron.

Installing:

	go get github.com/matipan/gobot && go install github.com/matipan/gobot/platforms/particle

Example:

	package main

	import (
		"time"

		"github.com/matipan/gobot"
		"github.com/matipan/gobot/drivers/gpio"
		"github.com/matipan/gobot/platforms/particle"
	)

	func main() {
		core := paticle.NewAdaptor("device_id", "access_token")
		led := gpio.NewLedDriver(core, "D7")

		work := func() {
			gobot.Every(1*time.Second, func() {
				led.Toggle()
			})
		}

		robot := gobot.NewRobot("particle",
			[]gobot.Connection{core},
			[]gobot.Device{led},
			work,
		)

		robot.Start()
	}

For further information refer to Particle readme:
https://github.com/hybridgroup/gobot/blob/master/platforms/particle/README.md
*/
package particle // import "github.com/matipan/gobot/platforms/particle"
