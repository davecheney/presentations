package main

func (m *Mux) SendMsg(msg string) {
	result := make(chan error, 1)
	m.ops <- func(m map[net.Addr]net.Conn) {
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
	conns := make(map[net.Addr]net.Conn)
	for _, op := range m.ops {
		op(conns)
	}
}
