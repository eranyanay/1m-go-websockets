package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"syscall"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

var EventsCollector *eventsCollector

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		return
	}
	if err := EventsCollector.Add(conn); err != nil {
		log.Printf("Failed to add connection %v", err)
		conn.Close()
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
			log.Fatalf("pprof failed: %v", err)
		}
	}()

	// Start event
	var err error
	EventsCollector, err = MkEventsCollector()
	if err != nil {
		panic(err)
	}

	go Start()

	http.HandleFunc("/", wsHandler)
	if err := http.ListenAndServe("0.0.0.0:8000", nil); err != nil {
		log.Fatal(err)
	}
}

func Start() {
	for {
		connections, err := EventsCollector.Wait()
		if err != nil {
			log.Printf("Failed to epoll wait %v", err)
			continue
		}
		for _, conn := range connections {
			if conn == nil {
				break
			}
			if _, _, err := wsutil.ReadClientData(conn); err != nil {
				if err := EventsCollector.Remove(conn); err != nil {
					log.Printf("Failed to remove %v", err)
				}
				conn.Close()
			} else {
				// This is commented out since in demo usage, stdout is showing messages sent from > 1M connections at very high rate
				//log.Printf("msg: %s", string(msg))
			}
		}
	}
}
