package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

func main() {
	gbot := gobot.NewGobot()
	api.NewAPI(gbot).Start()

	firmataAdaptor := firmata.NewFirmataAdaptor("arduino", "/dev/ttyACM0")

	led11 := gpio.NewLedDriver(firmataAdaptor, "led11", "11")
	led12 := gpio.NewLedDriver(firmataAdaptor, "led12", "12")
	led13 := gpio.NewLedDriver(firmataAdaptor, "led13", "13")

	robot := gobot.NewRobot("minions-bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led11, led12, led13},
		nil,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
