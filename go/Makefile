all: build run

build:
	go build -o arduino-minions

run: chown
	./arduino-minions

scan:
	gort scan serial

upload:
	gort arduino upload firmata /dev/ttyACM0

chown:
	@sudo chown ${USER}:${USER} /dev/ttyACM0

test:
	curl http://localhost:3000/api/robots/minions-bot/devices/led12/commands/Toggle