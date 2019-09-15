#!/usr/bin/env bash
#
# dev-docker.sh
# Copyright (C) 2019 jason <jason@t470slaptop>
#
# Distributed under terms of the MIT license.
#

docker-compose -f docker-compose-dev.yml down
docker-compose -f docker-compose-dev.yml build
docker-compose -f docker-compose-dev.yml up -d

exit 0

