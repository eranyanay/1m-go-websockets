#!/bin/bash

docker rm -vf $(docker ps -q --filter label=1m-go-websockets)
