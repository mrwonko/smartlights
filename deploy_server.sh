#!/usr/bin/env bash
set -eu -o pipefail

export GOOS=linux
export CGO_ENABLED=0

go build -o bin/lightserver ./lightserver/

host=mrwonko.de
ssh root@$host systemctl stop smartlights
scp bin/lightserver $host:/home/willi/
ssh root@$host systemctl start smartlights
