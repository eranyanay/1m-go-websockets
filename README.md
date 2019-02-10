# Going Infinite, handling 1M websockets connections in Go
Nothing fancy, this repository holds the complete implementation of the examples seen in Gophercon Israel talk, 2019.

# Usage
This repository demonstrates how a very high scale number of websockets connections can be maintained in Linux
Everything is written in pure Go
Each folder shows an example of a server implementation that overcomes various issues raised by the OS, by the hardware or the Go runtime itself,
as shown during the talk.

`setup.sh` is a wrapper to running multiple instances using Docker. See content of the script for more details of how to use it.
`destroy.sh` is a wrapper to stop all running clients. Note that it removes any running container, so use with caution.

Slides are available [here](https://speakerdeck.com/eranyanay/going-infinite-handling-1m-websockets-connections-in-go)
