# Going Infinite, handling 1M websockets connections in Go
This repository holds the complete implementation of the examples seen in Gophercon Israel talk, 2019.

> Going Infinite, handling 1 millions websockets connections in Go / Eran Yanay &mdash; [ [Video](https://www.youtube.com/watch?v=LI1YTFMi8W4) | [Slides](https://speakerdeck.com/eranyanay/going-infinite-handling-1m-websockets-connections-in-go) ]

It doesnt intend or claim to serve as a better, more optimal implementation than other libraries that implements the websocket protocol, it simply shows a set of tools, all combined together to demonstrate a server written in pure Go that is able to serve more than a million websockets connections with less than 1GB of ram.

# Usage
This repository demonstrates how a very high number of websockets connections can be maintained efficiently in Linux

Everything is written in pure Go

Each folder shows an example of a server implementation that overcomes various issues raised by the OS, by the hardware or the Go runtime itself, as shown during the talk.

`setup.sh` is a wrapper to running multiple instances using Docker. See content of the script for more details of how to use it.

`destroy.sh` is a wrapper to stop all running clients.

A single client instance can be executed by running `go run client.go -conn=<# connections to establish>`

# Remarks
This repo consists of a set of examples that were demonstrated during a live talk in Gophercon. 

What you see is what you get - while the implementation is fully functional, it doesn't intend to serve as a production-ready code, and no new features that were asked will be added. 

The only purpose of this repository is to serve as a reference and a case study.
