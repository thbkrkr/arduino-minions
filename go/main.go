package main

import (
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"

	"github.com/thbkrkr/qli/client"
)

func main() {
	gbot := gobot.NewGobot()

	/*	a := api.NewAPI(gbot)
		allowedOrigin := []string{"http://io:8000"}
		a.Start()
		a.AddHandler(AllowRequestsFrom(allowedOrigin[0]))
	*/

	firmataAdaptor := firmata.NewFirmataAdaptor("arduino", "/dev/ttyACM0")

	bigRed := gpio.NewLedDriver(firmataAdaptor, "bigR", "11")
	smallRed := gpio.NewLedDriver(firmataAdaptor, "smallRed", "12")
	smallGreen := gpio.NewLedDriver(firmataAdaptor, "smallGreen", "13")

	qli, err := client.NewClientFromEnv("arduinobot")
	if err != nil {
		panic(err)
	}

	work := func() {

		interval := time.Duration(2 * time.Second)
		alertTicker := time.NewTicker(interval).C
		sub, err := qli.Sub()
		if err != nil {
			panic(err)
		}

		lastMessage := time.Now()
		alert := false

		for {
			select {
			case <-alertTicker:
				if time.Since(lastMessage) > interval {
					alert = true
					smallGreen.Off()
					smallRed.On()
					bigRed.On()
				}
			case <-sub:
				if alert {
					smallRed.Off()
					bigRed.Off()
					alert = false
				}
				smallGreen.On()
				time.Sleep(100 * time.Millisecond)
				smallGreen.Off()
				lastMessage = time.Now()
			}
		}
	}

	robot := gobot.NewRobot("minions-bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{bigRed, smallRed, smallGreen},
		work,
	)

	gbot.AddRobot(robot)
	gbot.Start()
}
