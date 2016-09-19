package main

type Mux struct {
	conns map[net.Addr]net.Conn

	ops chan func(*Mux)
}

func (m *Mux) Add(conn net.Conn) {
	m.ops <- func(m *Mux) {
		m.conns[conn.RemoteAddr()] = conn
	}
}

func (m *Mux) Remove(addr net.Addr) {
	m.ops <- func(m *Mux) {
		delete(m.conns, addr)
	}

}

func (m *Mux) SendMsg(msg string) error {
	m.ops <- func(m *Mux) {
		for _, conn := range m.conns {
			io.WriteString(conn, msg)
		}
	}
	return nil
}

func (m *Mux) loop() {
	for _, op := range m.ops {
		op(m)
	}
}
