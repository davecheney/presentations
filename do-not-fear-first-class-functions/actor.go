package main

import (
	"io"
	"net"
)

type Mux struct {
	ops chan func(*Mux)
}

func (m *Mux) Add(conn net.Conn) {
	m.ops <- func(m map[net.Addr]net.Conn) {
		m[conn.RemoteAddr()] = conn
	}
}

func (m *Mux) Remove(addr net.Addr) {
	m.ops <- func(m map[net.Addr]net.Conn) {
		delete(m, addr)
	}
}

func (m *Mux) SendMsg(msg string) error {
	m.ops <- func(m map[net.Addr]net.Conn) {
		for _, conn := range m {
			io.WriteString(conn, msg)
		}
	}
	return nil
}

func (m *Mux) loop() {
	conns := make(map[net.Addr]net.Conn)
	for _, op := range m.ops {
		op(conns)
	}
}
