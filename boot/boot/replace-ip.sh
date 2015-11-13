#!/bin/bash -eu

HOME=/home/pi

ip=$(ifconfig eth0 | grep addr: | grep -o "[0-9\.]*" | head -1)

sed -i "s/localhost:/$ip:/" $HOME/arduino-minions/www/js/main.js
