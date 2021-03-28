#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error, print all commands.
set -v

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1
# docker-compose -f docker-compose-Orderer.yml down

docker-compose -f docker-compose-Orderer.yml up -d  #orderer.example.com 
docker ps -a

