// +build example
//
// Do not build by default.

package main

import (
	"fmt"

	"github.com/matipan/gobot"
	"github.com/matipan/gobot/drivers/aio"
	"github.com/matipan/gobot/drivers/i2c"
	"github.com/matipan/gobot/platforms/raspi"
)

func main() {
	board := raspi.NewAdaptor()
	gp := i2c.NewGrovePiDriver(board)
	sensor := aio.NewGroveRotaryDriver(gp, "A1")

	work := func() {
		sensor.On(aio.Data, func(data interface{}) {
			fmt.Println("sensor", data)
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{board},
		[]gobot.Device{gp, sensor},
		work,
	)

	robot.Start()
}
