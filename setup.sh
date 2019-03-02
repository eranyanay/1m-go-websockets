#!/bin/bash
## This script helps regenerating multiple client instances in different network namespaces using Docker
## This helps to overcome the ephemeral source port limitation
## Usage: ./connect <connections> <number of clients> <server ip>
## Server IP is usually the Docker gateway IP address, which is 172.17.0.1 by default
## Number of clients helps to speed up connections establishment at large scale, in order to make the demo faster

CONNECTIONS=$1
REPLICAS=$2
IP=$3
GOOS=linux go build --tags "static netgo" -o client client.go
for (( c=0; c<${REPLICAS}; c++ ))
do
    docker run -l 1m-go-websockets -v $(pwd)/client:/client -d alpine /client -conn=${CONNECTIONS} -ip=${IP}
done
