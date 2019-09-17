#!/usr/bin/env bash
#
# dev-run.sh
# Copyright (C) 2019 jason <jason@t470slaptop>
#
# Distributed under terms of the MIT license.
#

if [[ ! -e ./photos ]];then
	ansi --red  "Photos dir not present or sym-linked!"
	ansi --cyan "Run: ln -s ./client/static/photos"
	exit 1
fi

SERVER_PORT=:8089 \
SERVER_HOST=http://localhost \
PHOTO_PATH=photos \
go run server/server.go

exit 0

