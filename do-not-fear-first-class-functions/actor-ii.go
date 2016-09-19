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
	result := make(chan error, 1)
	m.ops <- func(m *Mux) {
		for _, conn := range m.conns {
			if err := io.WriteString(conn, msg); err != nil {
				result <- err
				return
			}
		}
		result <- nil
	}
	return <-result
}

func (m *Mux) loop() {
	for _, op := range m.ops {
		op(m)
	}
}
