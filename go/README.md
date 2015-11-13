# Play with Arduino

## Install Go and Gort

Install [go](https://golang.org/dl/).

Install [gort](http://gort.io/documentation/getting_started/downloads/).

Upload the firmdata to the Arduino with Gort.

	gort arduino install
	
	gort scan serial

	gort arduino upload firmata /dev/ttyACM0

## Install Gobot

	go get -d -u github.com/hybridgroup/gobot/...

## Build

	go build