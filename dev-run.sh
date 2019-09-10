#!/usr/bin/env bash
#
# dev-run.sh
# Copyright (C) 2019 jason <jason@t470slaptop>
#
# Distributed under terms of the MIT license.
#

SERVER_PORT=:8089 \
SERVER_HOST=http://localhost \
PHOTO_PATH=photos \
go run server.go

exit 0

