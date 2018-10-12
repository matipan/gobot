/*
Package keyboard contains the Gobot drivers for keyboard support.

Installing:

Then you can install the package with:

	go get github.com/matipan/gobot && go install github.com/matipan/gobot/platforms/keyboard

Example:

	package main

	import (
		"fmt"

		"github.com/matipan/gobot"
		"github.com/matipan/gobot/platforms/keyboard"
	)

	func main() {
		keys := keyboard.NewDriver()

		work := func() {
			keys.On(keyboard.Key, func(data interface{}) {
				key := data.(keyboard.KeyEvent)

				if key.Key == keyboard.A {
					fmt.Println("A pressed!")
				} else {
					fmt.Println("keyboard event!", key, key.Char)
				}
			})
		}

		robot := gobot.NewRobot("keyboardbot",
			[]gobot.Connection{},
			[]gobot.Device{keys},
			work,
		)

		robot.Start()
	}


For further information refer to keyboard README:
https://github.com/hybridgroup/gobot/blob/master/platforms/keyboard/README.md
*/
package keyboard // import "github.com/matipan/gobot/platforms/keyboard"
