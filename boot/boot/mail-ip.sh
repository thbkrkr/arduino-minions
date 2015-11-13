#!/bin/bash -eu

HERE=$(dirname "$(readlink -f "$0")")
MANDRILL_KEY="$(cat $HERE/.mandrill_key)"

mail() {
    local ip="$1"

    curl -A 'pi/0.1' \
        https://mandrillapp.com/api/1.0/messages/send.json \
        -d '{
      "key": "'$MANDRILL_KEY'",
      "message": {
        "html": "<h2>Hello dude! </h2> <p>This is my IP: <a href=\"http://'$ip':8000\">'$ip'</a></p>",
        "text": "",
        "subject": "[Pi] Connected on '"$ip"'",
        "from_email": "pi@blurb.space",
        "from_name": "Pi Pi",
        "to": [
          {
            "email": "thb.richard+pi@gmail.com",
            "type": "to"
          }
        ]
      }
    }'
}

get_ip() {
  ifconfig eth0 | grep addr: | grep -o "[0-9\.]*" | head -1
}

main() {
  local try=30
  local ip=""

  while [[ "$ip" == "" ]] || [[ $try -eq 0 ]]
  do
    ip=$(get_ip)
    sleep 1
    try=$(expr $try - 1)
  done

  [[ $try -eq 0 ]] && "mail ip_not_found" && exit 1

  mail $ip
}

main
