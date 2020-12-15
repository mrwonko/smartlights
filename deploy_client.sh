#!/usr/bin/env bash
set -eu -o pipefail

export GOOS=linux
export GOARCH=arm
export GOARM=5

echo "building"
go build -o bin/piclient ./piclient/

hosts=('raspi-light' 'raspi-light-kitchen')
hostsuffix=""
for host in "${hosts[@]}"; do
	echo "deploying to $host"
	ssh "$host$hostsuffix" sudo systemctl stop smartlights
	scp bin/piclient "$host:/home/pi/smartlights"
	ssh "$host$hostsuffix" sudo systemctl start smartlights
done
