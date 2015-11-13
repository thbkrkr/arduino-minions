#!/bin/bash

HOME=/home/pi
LOG=$HOME/logs
BOOT=$HOME/bin/boot

$BOOT/mail-ip.sh
$BOOT/replace-ip.sh

# Start Go binary for the Arduino communication
nohup $HOME/arduino-minions/go/arduino-minions > $LOG/arduino-minions-go.log &

# Start a Web Server
cd $HOME/arduino-minions/www && nohup python -m SimpleHTTPServer > $LOG/arduino-minions-www.log &
