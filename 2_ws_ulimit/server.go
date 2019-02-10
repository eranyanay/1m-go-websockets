package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync/atomic"
	"syscall"
)

var count int64

func ws(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	new := atomic.AddInt64(&count, 1)
	if new%100 == 0 {
		log.Printf("Total number of connections: %v", new)
	}
	defer func() {
		new := atomic.AddInt64(&count, -1)
		if new%100 == 0 {
			log.Printf("Total number of connections: %v", new)
		}
	}()

	// Read messages from socket
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		log.Printf("msg: %s", string(msg))
	}
}

func main() {
	// Increase resources limitations
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}

	// Enable pprof hooks
	go func() {
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			log.Fatalf("Pprof failed: %v", err)
		}
	}()

	http.HandleFunc("/", ws)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
