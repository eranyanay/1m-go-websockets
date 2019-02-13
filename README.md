# Going Infinite, handling 1M websockets connections in Go
This repository holds the complete implementation of the examples seen in Gophercon Israel talk, 2019.

# Usage
This repository demonstrates how a very high number of websockets connections can be maintained efficiently in Linux

Everything is written in pure Go

Each folder shows an example of a server implementation that overcomes various issues raised by the OS, by the hardware or the Go runtime itself, as shown during the talk.

`setup.sh` is a wrapper to running multiple instances using Docker. See content of the script for more details of how to use it.

`destroy.sh` is a wrapper to stop all running clients.

A single client instance can be executed by running `go run client.go -conn=<# connections to establish>`

Slides are available [here](https://speakerdeck.com/eranyanay/going-infinite-handling-1m-websockets-connections-in-go)
