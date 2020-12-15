#!/usr/bin/env bash
set -eu -o pipefail

export GOOS=linux

go build -o bin/lightserver ./lightserver/

host=mrwonko.de
ssh root@$host systemctl stop smartlights
scp bin/lightserver $host:/home/willi/
ssh root@$host systemctl start smartlights
