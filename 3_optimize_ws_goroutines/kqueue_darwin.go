// +build darwin

package main

import (
	"log"
	"reflect"
	"sync"
	"syscall"

	"github.com/gorilla/websocket"
)

type eventsCollector struct {
	fd          int
	connections map[int]*websocket.Conn
	// kqueue will watch these Kevent_t changes after call Kevent()
	// see more in freeBSD paper: https://people.freebsd.org/~jlemon/papers/kqueue.pdf
	changes []syscall.Kevent_t
	lock    *sync.RWMutex
}

func MkEventsCollector() (*eventsCollector, error) {
	fd, err := syscall.Kqueue()
	if err != nil {
		return nil, err
	}
	kevent := syscall.Kevent_t{
		Ident:  0,
		Filter: syscall.EVFILT_USER,
		Flags:  syscall.EV_ADD | syscall.EV_CLEAR,
	}
	if _, err = syscall.Kevent(fd, []syscall.Kevent_t{kevent}, nil, nil); err != nil {
		return nil, err
	}
	return &eventsCollector{
		fd:          fd,
		lock:        &sync.RWMutex{},
		connections: make(map[int]*websocket.Conn),
	}, nil
}

func (e *eventsCollector) Add(conn *websocket.Conn) error {
	fd := websocketFD(conn)
	e.changes = append(e.changes,
		syscall.Kevent_t{
			Ident: uint64(fd), Flags: syscall.EV_ADD, Filter: syscall.EVFILT_READ,
		},
	)
	e.lock.Lock()
	defer e.lock.Unlock()
	e.connections[fd] = conn
	if len(e.connections)%100 == 0 {
		log.Printf("Total number of connections: %v", len(e.connections))
	}
	return nil
}

func (e *eventsCollector) Remove(conn *websocket.Conn) error {
	fd := websocketFD(conn)
	e.changes = append(e.changes,
		syscall.Kevent_t{
			Ident: uint64(fd), Flags: syscall.EV_DELETE, Filter: syscall.EVFILT_READ,
		},
	)
	e.lock.Lock()
	defer e.lock.Unlock()
	e.connections[fd] = conn
	if len(e.connections)%100 == 0 {
		log.Printf("Total number of connections: %v", len(e.connections))
	}
	return nil
}

func (e *eventsCollector) Wait() ([]*websocket.Conn, error) {
	events := make([]syscall.Kevent_t, 100)
	n, err := syscall.Kevent(e.fd, e.changes, events, nil)
	if err != nil {
		return nil, err
	}
	e.lock.RLock()
	defer e.lock.RUnlock()
	var connections []*websocket.Conn
	for i := 0; i < n; i++ {
		conn := e.connections[int(events[i].Ident)]
		connections = append(connections, conn)
	}
	return connections, nil
}

func websocketFD(conn *websocket.Conn) int {
	connVal := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn").Elem()
	tcpConn := reflect.Indirect(connVal).FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")
	return int(pfdVal.FieldByName("Sysfd").Int())
}
