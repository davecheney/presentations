package main

import (
	"io"
	"net"
	"sync"
)

type Mux struct {
	mu    sync.Mutex
	conns map[net.Addr]net.Conn
}

func (m *Mux) Add(conn net.Conn) {
	mu.Lock()
	defer mu.Unlock()
	m.conns[conn.RemoteAddr()] = conn
}

func (m *Mux) Remove(addr net.Addr) {
	mu.Lock()
	defer mu.Unlock()
	delete(m.conns, addr)
}

func (m *Mux) SendMsg(msg string) error {
	mu.Lock()
	defer mu.Unlock()
	for _, conn := range m.conns {
		if err := io.WriteString(conn, msg); err != nil {
			return err
		}
	}
	return nil
}
